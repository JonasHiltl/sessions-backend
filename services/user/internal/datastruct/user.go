package datastruct

type PublicUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"              validate:"required"`
	FirstName string `json:"firstName"             validate:"required"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email"                 validate:"required,email"`
	Picture   string `json:"picture,omitempty"`
	Role      string `json:"role"`
}

type RequestUser struct {
	Username  string `json:"username"              validate:"required"`
	FirstName string `json:"firstame"              validate:"required"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email"                 validate:"required,email"`
	Password  string `json:"password,omitempty"    validate:"required,gte=8"`
	Picture   string `json:"picture,omitempty"`
}
