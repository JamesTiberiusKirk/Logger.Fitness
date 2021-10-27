package models

// LoginResponseDto struct that is getting returned to the user at login
type LoginResponseDto struct {
	Message string `json:"message"`
	Jwt     string `json:"jwt"`
}
