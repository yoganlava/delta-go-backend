package access

import (
	"context"
	"main/db"
	"main/internal/dto"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

func New() AccessService {
	return AccessService{db.Connection()}
}

type AccessService struct {
	pool *pgxpool.Pool
}

var AccessTypeEnum = map[string]int{
	"POST":   1,
	"PATCH":  1,
	"GET":    1,
	"DELETE": 1,
}

// func ModelAccess(setting entity.AccessCreatorSetting, accessModel string) (allowed bool) {

// 	return true
// }

func (as AccessService) MethodAccess(accessDTO dto.AccessDTO) (allowed bool) {
	var access entity.Access
	err := pgxscan.Get(context.Background(), as.pool, &access, `
	select a.name, a.id, a.setting
	from creator_access ua
	inner join access a on a.id = ua.access_id
	where ua.creator_id = $1 and ua.user_id = $2
`, accessDTO.AccessedID, accessDTO.UserID)
	if err != nil {
		return false
	}
	switch accessDTO.AccessModel {
	case "creator":
		switch accessDTO.Method {
		case "POST":
			return access.Setting.CreateCreator
		case "PATCH":
			return access.Setting.UpdateCreator
		case "DELETE":
			return access.Setting.DeleteCreator
		}
	case "project":
		switch accessDTO.Method {
		case "POST":
			return access.Setting.CreateProject
		case "PATCH":
			return access.Setting.UpdateProject
		case "DELETE":
			return access.Setting.DeleteProject
		}
	case "tier":
		switch accessDTO.Method {
		case "POST":
			return access.Setting.CreateTier
		case "PATCH":
			return access.Setting.UpdateTier
		case "DELETE":
			return access.Setting.DeleteTier
		}
	case "comment":
		switch accessDTO.Method {
		case "POST":
			return access.Setting.CreateComment
		case "PATCH":
			return access.Setting.UpdateComment
		case "DELETE":
			return access.Setting.DeleteComment
		}
	case "post":
		switch accessDTO.Method {
		case "POST":
			return access.Setting.CreatePost
		case "PATCH":
			return access.Setting.UpdatePost
		case "DELETE":
			return access.Setting.DeletePost
		}
	}

	return false
}
