package controller

type SignInData struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegistrationData struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
