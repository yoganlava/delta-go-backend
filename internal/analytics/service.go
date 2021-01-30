package analytics

import (
	"context"
	"main/db"
	"main/internal/dto"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AnalyticsService struct {
	pool *pgxpool.Pool
}

func New() AnalyticsService {
	return AnalyticsService{db.Connection()}
}

func (as AnalyticsService) GetProjectMonthlyEarnings(getProjectMonthlyEarningDTO dto.GetProjectMonthlyEarningsDTO) {

}

func (as AnalyticsService) GetProjectViews(getProjectViews dto.GetProjectViews) (int, error) {
	var creatorID int
	err := as.pool.QueryRow(context.Background(), `
	select id as creatorID from creator
	where user_id = $1
	`,
		getProjectViews.UserID,
	).Scan(&creatorID)
	if err != nil {
		return 0, err
	}

	var views int
	err = as.pool.QueryRow(context.Background(), `
	select views from total_project_view
	where project_id = $1
	`,
		getProjectViews.ProjectID,
	).Scan(&views)
	if err != nil {
		return 0, err
	}
	return views, err
}

// func (as AnalyticsService) GetLastPayouts(getLastPayouts dto.GetLastPayoutsDTO) error {
// 	var creatorID int
// 	err := as.pool.QueryRow(context.Background(), `
// 	select id as creatorID from creator
// 	where user_id = $1
// 	`,
// 		getLastPayouts.UserID,
// 	).Scan(&creatorID)
// 	if err != nil {
// 		return err
// 	}
// 	var payouts []*entity.PayoutMonth
// 	err = pgxscan.Select(context.Background(), as.pool, &payouts, `
// 	select t.created_at as date, t.amount from
// 	`)
// }
