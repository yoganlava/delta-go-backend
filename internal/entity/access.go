package entity

type Access struct {
	ID      int           `json:"id"`
	Site    bool          `json:"site"`
	Name    string        `json:"name"`
	Setting AccessSetting `json:"setting"`
}

type AccessSetting struct {
	CreateCreator  bool `json:"create_creator"`
	UpdateCreator  bool `json:"update_creator"`
	DeleteCreator  bool `json:"delete_creator"`
	CreateProject  bool `json:"create_project"`
	UpdateProject  bool `json:"update_project"`
	DeleteProject  bool `json:"delete_project"`
	CreateTier     bool `json:"create_tier"`
	UpdateTier     bool `json:"update_tier"`
	DeleteTier     bool `json:"delete_tier"`
	CreatePost     bool `json:"create_post"`
	UpdatePost     bool `json:"update_post"`
	DeletePost     bool `json:"delete_post"`
	ContactCreator bool `json:"contact_creator"`
	UpdateComment  bool `json:"update_comment"`
	DeleteComment  bool `json:"delete_comment"`
	CreateComment  bool `json:"create_comment"`
	// ProjectSetting AccessProjectSetting `json:"project_setting"`
	// CreatorSetting AccessCreatorSetting `json:"creator_setting"`
	// SiteSetting    AccessSiteSetting    `json:"site_setting"`
}

type AccessProjectSetting struct {
	CreatePost    bool `json:"create_post"`
	UpdatePost    bool `json:"update_post"`
	DeletePost    bool `json:"delete_post"`
	UpdateProject bool `json:"update_project"`
	DeleteProject bool `json:"delete_project"`
	CreateTier    bool `json:"create_tier"`
	UpdateTier    bool `json:"update_tier"`
	DeleteTier    bool `json:"delete_tier"`
}

type AccessCreatorSetting struct {
	CreateCreator bool `json:"create_creator"`
	UpdateCreator bool `json:"update_creator"`
	DeleteCreator bool `json:"delete_creator"`
	CreateProject bool `json:"create_project"`
	UpdateProject bool `json:"update_project"`
	DeleteProject bool `json:"delete_project"`
	CreateTier    bool `json:"create_tier"`
	UpdateTier    bool `json:"update_tier"`
	DeleteTier    bool `json:"delete_tier"`
}

type AccessSiteSetting struct {
	UpdateCreator bool `json:"update_creator"`
	DeleteCreator bool `json:"delete_creator"`
	UpdatePost    bool `json:"update_post"`
	DeletePost    bool `json:"delete_post"`
	UpdateProject bool `json:"update_project"`
	DeleteProject bool `json:"delete_project"`
}
