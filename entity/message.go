package entity

type ErrorMessage struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
