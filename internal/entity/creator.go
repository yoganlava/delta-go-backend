package entity

import "time"

//Creator model
type Creator struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Bio           string    `json:"bio"`
	AvatarImageID int       `json:"avatar_image_id"`
	UserID        int       `json:"user_id"`
	CreatorRankID int       `json:"creator_rank_id"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}
