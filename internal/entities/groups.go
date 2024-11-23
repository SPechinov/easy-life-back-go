package entities

type Group struct {
	ID        string
	Name      string
	IsPayed   bool
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
}

type GroupFull struct {
	Group
	Users []GroupUser
}

type GroupsGetList struct {
	UserID string
}

type GroupGet struct {
	ID string
}

type GroupGetInfo struct {
	ID     string
	UserID string
}

type GroupAdd struct {
	Name    string
	AdminID string
}

type GroupPatch struct {
	ID     string
	UserID string
	Name   *string
}

type GroupDelete struct {
	ID     string
	UserID string
}

type GroupDeleteConfirm struct {
	ID     string
	UserID string
	Code   string
}
