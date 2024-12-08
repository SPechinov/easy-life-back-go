package entities

type User struct {
	ID        string
	Email     string
	Phone     string
	Password  string
	FirstName string
	LastName  *string
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
}

type UserGet struct {
	ID    string
	Email string
	Phone string
}

type UserAuthWay struct {
	Email string
	Phone string
}

func (way *UserAuthWay) GetAuthValue() string {
	if way.Email != "" {
		return way.Email
	}

	return way.Phone
}

type UserLogin struct {
	AuthWay  UserAuthWay
	Password string
}

type UserAdd struct {
	AuthWay UserAuthWay
}

type UserAddConfirm struct {
	AuthWay   UserAuthWay
	FirstName string
	Password  string
	Code      string
}

type UserForgotPassword struct {
	AuthWay UserAuthWay
}

type UserForgotPasswordConfirm struct {
	AuthWay  UserAuthWay
	Password string
	Code     string
}

type UserUpdateJWT struct {
	ID         string
	SessionID  string
	RefreshJWT string
}

type UserLogout struct {
	ID        string
	SessionID string
}

type UserLogoutAll struct {
	ID string
}
