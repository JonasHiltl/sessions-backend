package dto

import "github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"

type User struct {
	Id        string
	Provider  datastruct.Provider
	Username  string
	Email     string
	Firstname string
	Lastname  string
	Password  string
	Avatar    string
}
