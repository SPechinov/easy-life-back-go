package entities

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

type GroupGetUsersList struct {
	GroupID string
}

type GroupInviteUser struct {
	GroupID string
	UserID  string
}

type GroupExcludeUser struct {
	GroupID string
	UserID  string
}
