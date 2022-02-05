package datastruct

type MessageRes struct {
	Message string `json:"message" validate:"required"`
}
