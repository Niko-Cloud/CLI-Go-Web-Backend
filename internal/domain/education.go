package domain

import "time"

type Education struct {
	ID          int        `json:"id"`
	Institution string     `json:"institution"`
	Degree      string     `json:"degree"`
	Major       *string    `json:"major,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description *string    `json:"description,omitempty"`
}
