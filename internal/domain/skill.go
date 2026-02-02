package domain

type Skill struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Level    *string `json:"level,omitempty"`
	Category *string `json:"category,omitempty"`
}
