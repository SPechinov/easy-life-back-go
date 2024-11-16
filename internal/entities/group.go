package entities

type Group struct {
	ID        string
	Name      string
	Admin     User
	IsPayed   bool
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
}

func (g Group) Deleted() bool {
	return g.DeletedAt != nil
}

type GroupAdd struct {
	Name    string
	AdminID string
}
