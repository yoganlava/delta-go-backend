package entity

import "time"

type Goal struct {
	ID          int       `json:"id"`
	Description *string   `json:"description"`
	Goal        float32   `json:"goal"`
	Type        string    `json:"type"`
	ProjectID   int       `json:"project_id"`
	CreatedAt   time.Time `json:"created_at`
	UpdatedAt   time.Time `json:"updated_at"`
}
