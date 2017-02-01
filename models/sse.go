package models

type (
	Sse struct {
		UserID string   `json:"u"`
		Val    float32  `json:"v"`
		Tags   []string `json:"t"`
	}
)
