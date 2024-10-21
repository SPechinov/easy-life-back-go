package redis

func GetKeyUserRegistrationCode(email string) string {
	return "http:users:reg-code:" + email
}

func GetKeyUsersForgotPasswordCode(email string) string {
	return "http:users:forgot-password-code:" + email
}

func GetKeyUserJWTPair(email string) string {
	return "http:users:auth:jwt-pair:" + email
}
