package entity

//Creator model
type Creator struct {
	ID            int
	Name          string
	Bio           string
	BannerImageID int    `db:"banner_image_id"`
	CoverImageID  int    `db:"cover_image_id"`
	UserID        int    `db:"user_id"`
	CreatorRankID int    `db:"creator_rank_id"`
	UpdatedAt     string `db:"updated_at"`
	CreatedAt     string `db:"created_at"`
}
