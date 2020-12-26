package dto

type CreateCreatorDTO struct {
	Name          string  `json:"name" binding:""`
	Bio           *string `json:"bio"`
	AvatarImageID *int    `json:"avatar_image_id"`
	UserID        int     `json:"user_id"`
}

type UpdateCreatorDTO struct {
	ID            int     `json:"id" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Bio           *string `json:"bio"`
	AvatarImageID *int    `json:"avatar_image_id"`
	CreatorRankID int     `json:"creator_rank_id"`
	UserID        int     `json:"user_id"`
}

type SearchCreatorDTO struct {
	Query string `json:"query"`
	*PaginationDTO
}
