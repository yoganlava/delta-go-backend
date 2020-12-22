package dto

type AccessMethod string

type AccessModel string

//login strategy enum definition
const (
	POST   AccessMethod = "POST"
	GET                 = "GET"
	DELETE              = "DELETE"
	PATCH               = "PATCH"
)

const (
	PROJECT AccessModel = "project"
	POSTM               = "post"
	TIER                = "tier"
	COMMENT             = "comment"
	CREATOR             = "creator"
)

type AccessDTO struct {
	UserID      int          `json:"user_id"`
	Method      AccessMethod `json:"method"`
	AccessModel AccessModel  `json:"access_model"`
	AccessedID  string       `json:"accessed_id"`
}
