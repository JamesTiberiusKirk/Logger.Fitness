package types

import (
	"gorm.io/gorm"
)

// User is a direct representation of the mongo document.
type User struct {
	gorm.Model
	ID             string            `json:"id"`
	Email          string            `json:"email"`
	Username       string            `json:"username"`
	Password       string            `json:"password"`
	Active         bool              `json:"active"`
	Roles          map[string]string `json:"roles"`
	ProfilePicture string            `json:"profile_picture"`
	OauthProvider  *OauthProvider    `json:"oauth_provider,omitempty"`
}

type OauthProvider struct {
	gorm.Model
	ID         string `json:"id" gorm:"primaryKey"`
	UserID     string `json:"user_id" gorm:"foreignKey:UserID"`
	Provider   string `json:"provider"`
	ProviderID string `json:"provider_id"`
}

// UserLoginForm is the input from user on login.
type UserLoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
