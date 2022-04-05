package comtypes

type MessageRes struct {
	Message string `json:"message" validate:"required"`
}

type JwtPayload struct {
	Iss  string `json:"iss"`
	Sub  string `json:"sub"`
	Iat  int    `json:"iat"`
	Role Role   `json:"role"`
}
