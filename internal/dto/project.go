package dto

type CreateProjectDTO struct {
	Name          string `json:"name"`
	PageURL       string `json:"page_url"`
	Description   string `json:"description"`
	Creating      string `json:"creating"`
	BannerImageID int
	CoverImageID  int
	CreatorID     int
	CoverID       int
	CategoryID    int
	Setting       string
}

type DonateToProjectDTO struct {
	ProjectID string `json:"project_id"`
}
