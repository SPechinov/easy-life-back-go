package group

type AddDTO struct {
	Name string `json:"name" form:"name"`
}

type PatchDTO struct {
	Name *string `json:"name" form:"name"`
}
