package entities

type Group struct {
	ID        string
	Name      string
	IsPayed   bool
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
}

func (g *Group) Deleted() bool {
	return g.DeletedAt != nil
}

type GroupFull struct {
	Group
	Users []GroupUser
}

type GroupsGetList struct {
	UserID  string
	Deleted bool
}

type GroupGet struct {
	ID string
}

type GroupGetInfo struct {
	ID string
}

type GroupAdd struct {
	Name    string
	AdminID string
}

type GroupPatch struct {
	ID   string
	Name *string
}

type GroupDelete struct {
	ID string
}
