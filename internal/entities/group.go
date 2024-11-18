package entities

type Group struct {
	ID        string
	Name      string
	Admin     GroupUser
	IsPayed   bool
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
	Users     []GroupUser
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

func (g Group) Deleted() bool {
	return g.DeletedAt != nil
}

type GroupGet struct {
	GroupID string
}

type GroupGetUsersList struct {
	GroupID string
}

type GroupAdd struct {
	Name    string
	AdminID string
}

type GroupPatch struct {
	GroupID string
	Name    *string
}

type GroupInviteUser struct {
	GroupID string
	UserID  string
}

type GroupExcludeUser struct {
	GroupID string
	UserID  string
}
