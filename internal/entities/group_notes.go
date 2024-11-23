package entities

type Note struct {
	ID          string
	Title       string
	Info        *NoteInfo
	GroupID     string
	UserCreator User
	UserUpdater User
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

type NoteInfo struct {
	Description string
}

type NoteAdd struct {
	UserID      string
	GroupID     string
	Title       string
	Description string
}

type NotePatch struct {
	ID          string
	UserID      string
	GroupID     string
	Title       *string
	Description *string
}

type NoteGetList struct {
	UserID  string
	GroupID string
	Deleted bool
}

type NoteGet struct {
	ID      string
	UserID  string
	GroupID string
}

type NoteDelete struct {
	ID      string
	UserID  string
	GroupID string
}
