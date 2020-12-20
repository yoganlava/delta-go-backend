package dto

type CreateCreatorDTO struct {
	Name          string  `json:"name" binding:""`
	Bio           *string `json:"bio"`
	AvatarImageID *int    `json:"avatar_image_id"`
	UserID        int     `json:"user_id"`
}
