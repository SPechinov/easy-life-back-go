package group_users

type InviteUserDTO struct {
	UserID string `json:"userId" form:"userId"`
}

type ExcludeUserDTO struct {
	UserID string `json:"userId" form:"userId"`
}
