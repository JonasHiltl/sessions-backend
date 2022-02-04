package datastruct

type LoginBody struct {
	UsernameOrEmail string `json:"usernameOrEmail" validate:"required"`
	Password        string `json:"password"        validate:"required"`
}

type AuthRes struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
