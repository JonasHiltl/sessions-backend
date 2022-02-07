package comtypes

type MessageRes struct {
	Message string `json:"message" validate:"required"`
}
