package models

type Pyq struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Stream string `json:"stream"`
	Year   string `json:"year"`
	Batch  string `json:"batch"`
}
