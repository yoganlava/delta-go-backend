package entity

import "time"

type Post struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Mature       bool      `json:"mature"`
	ProjectID    int       `json:"project_id"`
	SubmitUserID int       `json:"submit_user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PostText struct {
	PostID    int `json:"post_id"`
	WordCount int `json:"word_count"`
	CharCount int `json:"char_count"`
}

type PostFile struct {
	PostID int `json:"post_id"`
	FileID int `json:"file_id"`
}

type PostTag struct {
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id"`
}

type PostVote struct {
	UserID    int       `json:"user_id"`
	Vote      int       `json:"vote"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PostTier struct {
	PostID int `json:"post_id"`
	TierID int `json:"tier_id"`
}
