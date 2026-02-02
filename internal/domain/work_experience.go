package domain

import "time"

type WorkExperience struct {
	ID          int        `json:"id"`
	Company     string     `json:"company"`
	Position    string     `json:"position"`
	StartDate   time.Time  `json:"startDate"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description []string   `json:"description,omitempty"`
	TechStack   []string   `json:"techStack,omitempty"`
}
