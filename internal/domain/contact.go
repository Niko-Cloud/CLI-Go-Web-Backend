package domain

type Contact struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
