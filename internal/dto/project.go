package dto

type CreateProjectDTO struct {
	Name          string
	PageURL       string
	Description   string
	Creating      string
	BannerImageID int
	CoverImageID  int
	CreatorID     int
	CoverID       int
	CategoryID    int
	Setting       string
}
