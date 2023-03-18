package youdao

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Play(word string) (string, error) {
	resp, err := http.Get("https://dict.youdao.com/dictvoice?type=0&audio=" + word)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ct := resp.Header.Get("Content-Type")
	return fmt.Sprintf("data:%s;base64,%s", ct, base64.StdEncoding.EncodeToString(b)), nil
}
