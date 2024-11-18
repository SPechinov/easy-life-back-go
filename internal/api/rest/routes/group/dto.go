package group

type AddDTO struct {
	Name string `json:"name" form:"name"`
}

type PatchDTO struct {
	Name *string `json:"name" form:"name"`
}

type InviteDTO struct {
	UserID string `json:"userId" form:"userId"`
}

type ExcludeDTO struct {
	UserID string `json:"userId" form:"userId"`
}
