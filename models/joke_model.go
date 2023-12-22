package models

type Joke struct {
	Text     string `json:"text"`
	Type     string `json:"type"`
	Strength int    `json:"strength"`
}
