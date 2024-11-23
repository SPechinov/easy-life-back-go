package group_notes

type AddDTO struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}

type PatchDTO struct {
	Title       *string `json:"title" form:"title"`
	Description *string `json:"description" form:"description"`
}
