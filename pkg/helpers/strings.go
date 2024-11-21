package helpers

func StrToPtr(s string) *string {
	return &s
}

func BoolToPtr(value bool) *bool {
	return &value
}
