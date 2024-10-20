package redis

func GetKeyHttpUserRegistrationCode(email string) string {
	return "http:users:reg-code:" + email
}

func GetKeyHttpUsersForgotPasswordCode(email string) string {
	return "http:users:forgot-password-code:" + email
}

func GetKeyHttpUserJWTPair(email string) string {
	return "http:users:auth:jwt-pair:" + email
}
