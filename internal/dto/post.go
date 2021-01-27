package dto

type FetchProjectPostsDTO struct {
	ProjectID int    `form:"project_id"`
	OrderBy   string `form:"order_by"`
	Limit     int    `form:"limit"`
	Page      int    `form:"page"`
	Type      string `form:"type"`
	Mature    int    `form:"mature"`
}

type DeleteProjectPostDTO struct {
	PostID int `json:"post_id"`
	UserID int
}
