package dto

type GetProjectMonthlyEarningsDTO struct {
	ProjectID int `json:"project_id"`
	UserID    int
}

type GetLastPayoutsDTO struct {
	ProjectID int `json:"project_id"`
	UserID    int
}

type GetProjectViews struct {
	ProjectID int `json:"project_id"`
	UserID    int
}
