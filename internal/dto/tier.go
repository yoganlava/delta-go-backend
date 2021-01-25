package dto

type CreateTierDTO struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	CoverImageID int     `json:"cover_image_id"`
	Price        float32 `json:"price"`
	ProjectID    int     `json:"project_id"`
}

type CreateTierBenefitDTO struct {
	Description   string `json:"description"`
	BenefitPeriod string `json:"benefit_period"`
}

type LinkTierBenefitDTO struct {
	TierID    int `json:"tier_id"`
	BenefitID int `json:"benefit_id"`
}

type EditTierBenefitDTO struct {
	CreateTierBenefitDTO
	BenefitID int `json:"benefit_id"`
}

type EditTierDTO struct {
	TierID       int     `json:"tier_id"`
	Description  string  `json:"description"`
	CoverImageID int     `json:"cover_image_id"`
	Price        float32 `json:"price"`
}
