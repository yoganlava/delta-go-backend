package entity

type Post struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Mature       bool   `json:"mature"`
	ProjectID    int    `json:"project_id"`
	SubmitUserID int    `json:"submit_user_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
