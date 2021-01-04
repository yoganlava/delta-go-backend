package entity

import "time"

type Tier struct {
	ID           int
	Title        string
	Description  string
	CoverImageID int
	Price        float32
	ProjectID    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
