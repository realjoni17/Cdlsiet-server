package models

type Books struct {
	ID      uint   `json:"id"`
	Stream  string `json:"stream"`
	Year    string `json:"year"`
	Url     string `json:"url"`
	Subject string `json:"subject"`
}
