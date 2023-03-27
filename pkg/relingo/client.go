package relingo

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/samber/lo"
)

type Client struct {
	cfg *Config

	hc *resty.Client
}

func NewClient(cfg *Config) *Client {
	c := &Client{
		cfg: cfg,
	}
	c.hc = resty.New().SetBaseURL(cfg.BaseURL).OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		if !c.Ready() {
			return errors.New("empty token")
		}

		return nil
	})
	return c
}

func (c *Client) SetToken(token string) {
	c.cfg.Token = token
}

func (c *Client) Ready() bool {
	return c.cfg.Token != ""
}

func (c *Client) GetUserInfo() (*RespUserInfo, error) {
	var userInfo RespUserInfo
	result := NewResponse(userInfo)
	resp, err := c.hc.R().SetHeaders(c.relingoHeaders()).SetResult(result).Post("/getUserInfo")
	if err != nil {
		return nil, err
	}

	if err := checkResult(resp, result); err != nil {
		return nil, err
	}

	return &result.Data, nil
}

// GetVocabularyList 获取生词本列表
func (c *Client) GetVocabularyList() ([]VocabularyListItem, error) {
	var data []VocabularyListItem
	result := NewResponse(data)
	resp, err := c.hc.R().SetHeaders(c.relingoHeaders()).SetResult(result).Post("/getVocabularyList")
	if err != nil {
		return nil, err
	}

	if err := checkResult(resp, result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetVocabulary 获取某个生词本的所有单词
func (c *Client) GetVocabulary(id, typ string) ([]string, error) {
	var data Vocabulary
	result := NewResponse(data)
	resp, err := c.hc.R().SetHeaders(c.relingoHeaders()).SetBody(&VocabularyBody{Id: id, Type: typ}).
		SetResult(result).Post("/getVocabulary")
	if err != nil {
		return nil, err
	}

	if err := checkResult(resp, result); err != nil {
		return nil, err
	}

	return result.Data.Words, nil
}

// MasteredWords 获取已掌握的单词
func (c *Client) MasteredWords(id string) ([]string, error) {
	return c.GetVocabulary(id, "mastered")
}

func (c *Client) SubmitVocabulary(words []string) error {
	vs, err := c.GetVocabularyList()
	if err != nil {
		return err
	}

	masteredItem, ok := lo.Find(vs, func(item VocabularyListItem) bool { return item.Type == "mastered" })
	if !ok {
		return fmt.Errorf("not found mastered uid")
	}

	var data Vocabulary
	result := NewResponse(data)
	resp, err := c.hc.R().SetHeaders(c.relingoHeaders()).SetBody(&VocabularyBody{Id: masteredItem.Id, Type: "mastered", Words: words}).
		SetResult(result).Post("/submitVocabulary")
	if err != nil {
		return err
	}

	return checkResult(resp, result)
}

func (c *Client) LockupDict(word string) (*DictItem, error) {
	data := make([]DictItem, 0)
	result := NewResponse(data)
	resp, err := c.hc.R().SetHeaders(c.relingoHeaders()).SetBody(&LockupBody{Text: word, To: "zh"}).SetResult(result).Post("/lookupDict")
	if err != nil {
		return nil, err
	} else if err := checkResult(resp, result); err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, errors.New("not found")
	}

	v := result.Data[0]
	return &v, nil
}

func (c *Client) relingoHeaders() map[string]string {
	return map[string]string{
		"User-Agent":        "relingo-desktop",
		"x-relingo-lang":    "cn",
		"x-relingo-token":   c.cfg.Token,
		"x-relingo-version": "2.5.0",
	}
}

func checkResult[Data RD](resp *resty.Response, result *Response[Data]) error {
	if resp.IsError() {
		return fmt.Errorf("call failed: %s", resp.String())
	}

	if result.Code != 0 {
		return fmt.Errorf("call relingo api: %s", result.Message)
	}

	return nil
}
