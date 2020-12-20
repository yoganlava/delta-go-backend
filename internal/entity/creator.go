package entity

import "time"

//Creator model
type Creator struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	AvatarImageID  int            `json:"avatar_image_id"`
	UserID         int            `json:"user_id"`
	CreatorRankID  int            `json:"creator_rank_id"`
	UpdatedAt      time.Time      `json:"updated_at"`
	CreatedAt      time.Time      `json:"created_at"`
	CreatorRank    CreatorRank    `json:"creator_rank"`
	CreatorProfile CreatorProfile `json:"creator_profile"`
}

type CreatorRank struct {
	ID            int     `json:"id"`
	Fee           float32 `json:"fee"`
	Name          string  `json:"name"`
	Importantance int     `json:"importance"`
}

type CreatorProfile struct {
	LegalFirstName string `json:"legal_first_name"`
	LegalLastName  string `json:"legal_last_name"`
	Bio            string `json:"bio"`
	CreatorID      int    `json:"creator_id"`
}
