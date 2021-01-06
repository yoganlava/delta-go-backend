package dto

import "main/internal/entity"

type CreateProjectDTO struct {
	Name           string `json:"name"`
	PageURL        string `json:"page_url"`
	Description    string `json:"description"`
	Creating       string `json:"creating"`
	BannerImageID  int    `json:"banner_image_id"`
	CoverImageID   int    `json:"cover_image_id"`
	CreatorID      int    `json:"creator_id"`
	CoverID        int    `json:"cover_id"`
	CategoryID     int    `json:"category_id"`
	IsCrowdFunding bool   `json:"is_crowd_funding"`
}

type DonateToProjectDTO struct {
	ProjectID string `json:"project_id"`
}

type FeedProjectDTO struct {
	Post         *entity.Post                    `json:"post"`
	Donation     *entity.DonationTransaction     `json:"donation"`
	Subscription *entity.SubscriptionTransaction `json:"subscription"`
	Type         string                          `json:"type"`
}
