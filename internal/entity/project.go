package entity

type Project struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	PageURL        string         `json:"page_url"`
	Description    string         `json:"description"`
	Creating       string         `json:"creating"`
	BannerImageID  int            `json:"banner_image_id"`
	CoverImageID   int            `json:"cover_image_id"`
	CreatorID      int            `json:"creator_id"`
	CoverID        int            `json:"cover_id"`
	CategoryID     int            `json:"category_id"`
	IsCrowdFunding bool           `json:"is_crowd_funding"`
	Setting        ProjectSetting `json:"setting"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
}

type ProjectSetting struct {
	ShowTotalEarning bool `json:"show_total_earning"`
	ShowTopSupporter bool `json:"show_top_supporter"`
	AnyOneCanComment bool `json:"anyone_can_comment"`
}

type CreatorProject struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	PageURL string `json:"page_url"`
}
