package dto

type CreateTierDTO struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	CoverImageID int     `json:"cover_image_id"`
	Price        float32 `json:"price"`
	ProjectID    int     `json:"project_id"`
}
