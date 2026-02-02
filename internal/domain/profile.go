package domain

type Profile struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Role        string  `json:"role"`
	Description string  `json:"description"`
	Location    *string `json:"location,omitempty"`
}
