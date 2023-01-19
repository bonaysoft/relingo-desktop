package relingo

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    struct {
		Words []Word `json:"words"`
	} `json:"data"`
}

type Word struct {
	Phonetic     []string `json:"phonetic"`
	Variant      []string `json:"variant"`
	Id           string   `json:"_id"`
	Source       string   `json:"source"`
	Display      *string  `json:"display,omitempty"`
	Translations []struct {
		Target string  `json:"target"`
		Pos    string  `json:"pos"`
		Score  float64 `json:"score"`
	} `json:"translations"`
	Lang      string `json:"lang"`
	Mastered  bool   `json:"mastered"`
	Stared    bool   `json:"stared"`
	Sentences []struct {
		Id        string `json:"_id"`
		Uid       string `json:"uid"`
		Word      string `json:"word"`
		Sentence  string `json:"sentence"`
		Url       string `json:"url"`
		CreatedAt int64  `json:"createdAt"`
		UpdatedAt int64  `json:"updatedAt"`
		Md5       string `json:"md5"`
	} `json:"sentences"`
	Revision   bool   `json:"revision"`
	NeedRevise bool   `json:"needRevise"`
	Scope      string `json:"scope,omitempty"`
}
