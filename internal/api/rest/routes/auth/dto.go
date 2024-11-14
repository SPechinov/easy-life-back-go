package auth

type LoginDTO struct {
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

type RegistrationDTO struct {
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
}

type RegistrationConfirmDTO struct {
	Email     string `json:"email" form:"email"`
	Phone     string `json:"phone" form:"phone"`
	FirstName string `json:"firstName" form:"firstName"`
	Password  string `json:"password" form:"password"`
	Code      string `json:"code" form:"code"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
}

type ForgotPasswordConfirmDTO struct {
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	Code     string `json:"code" form:"code"`
}
