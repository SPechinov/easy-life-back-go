package group

type AddDTO struct {
	Name string `json:"name" form:"name"`
}

type PatchDTO struct {
	Name *string `json:"name" form:"name"`
}

type InviteUserDTO struct {
	UserID string `json:"userId" form:"userId"`
}

type ExcludeUserDTO struct {
	UserID string `json:"userId" form:"userId"`
}

type DeleteConfirmDTO struct {
	Code string `json:"code" form:"code"`
}
