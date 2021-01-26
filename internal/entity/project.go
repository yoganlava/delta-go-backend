package entity

import "time"

type Project struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	PageURL        string         `json:"page_url"`
	Description    string         `json:"description"`
	Creating       string         `json:"creating"`
	Banner         string         `json:"banner"`
	Avatar         string         `json:"avatar"`
	CreatorID      *int           `json:"creator_id,omitempty"`
	CategoryID     *int           `json:"category_id,omitempty"`
	IsCrowdFunding *bool          `json:"is_crowd_funding,omitempty"`
	Setting        ProjectSetting `json:"setting"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at,omitempty"`
	Creator        CreatorProject `json:"creator"`
}

type SearchProject struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Creating string `json:"creating"`
	Category string `json:"category"`
}

type ProjectSetting struct {
	ShowTotalEarning bool `json:"show_total_earning,omitempty"`
	ShowTopSupporter bool `json:"show_top_supporter,omitempty"`
	AnyOneCanComment bool `json:"anyone_can_comment,omitempty"`
}

type CreatorProject struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsCompany bool   `json:"is_company"`
}

type Feed struct {
	Post         *Post                    `json:"post"`
	Donation     *DonationTransaction     `json:"donation"`
	Subscription *SubscriptionTransaction `json:"subscription"`
	Type         string                   `json:"type"`
}
