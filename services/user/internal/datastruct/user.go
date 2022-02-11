package datastruct

type PublicUser struct {
	ID          string `json:"id"`
	Username    string `json:"username"              validate:"required"`
	FirstName   string `json:"firstname"             validate:"required"`
	LastName    string `json:"lastname,omitempty"`
	Picture     string `json:"picture,omitempty"`
	Role        string `json:"role"`
	FriendCount int    `json:"friendCount,omitempty"`
}

type RequestUser struct {
	Username  string `json:"username"              validate:"required"`
	FirstName string `json:"firstname"              validate:"required"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email"                 validate:"required,email"`
	Password  string `json:"password,omitempty"    validate:"required,gte=8"`
	Picture   string `json:"picture,omitempty"`
}

func (pu PublicUser) AddCount(count int) PublicUser {
	return PublicUser{
		ID:          pu.ID,
		Username:    pu.Username,
		FirstName:   pu.FirstName,
		LastName:    pu.LastName,
		Picture:     pu.Picture,
		Role:        string(pu.Role),
		FriendCount: count,
	}
}
