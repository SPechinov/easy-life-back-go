package response

type Success struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

func NewSuccess(data interface{}) *Success {
	return &Success{
		Ok:   true,
		Data: data,
	}
}
