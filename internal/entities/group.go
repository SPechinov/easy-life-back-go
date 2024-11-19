package entities

type GroupInfo struct {
	ID        string
	Name      string
	IsPayed   bool
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
}

func (g GroupInfo) Deleted() bool {
	return g.DeletedAt != nil
}

type GroupFull struct {
	GroupInfo
	Users []GroupUser
}

type GroupUser struct {
	ID         string
	Email      string
	Phone      string
	FirstName  string
	LastName   *string
	Permission int
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  *string
	InvitedAt  string
}

type GroupsGetList struct {
	UserID string
}

type GroupGet struct {
	ID string
}

type GroupGetInfo struct {
	ID string
}

type GroupGetUsersList struct {
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

type GroupInviteUser struct {
	ID     string
	UserID string
}

type GroupExcludeUser struct {
	ID     string
	UserID string
}
