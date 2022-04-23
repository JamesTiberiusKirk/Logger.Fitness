package types

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is a direct representation of the mongo document.
type User struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Email          string             `json:"email" bson:"email" validate:"required,email"`
	Username       string             `json:"username" bson:"username" validate:"required,min=3,max=100"`
	Password       string             `json:"password" bson:"password" validate:"required,min=7,max=255"`
	Active         bool               `bson:"active"`
	Roles          map[string]string  `json:"roles,omitempty" bson:"roles"`
	ProfilePicture string             `json:"profile_picture,omitempty" bson:"profile_picture"`
	Provider       string             `bson:"provider"`
	ProviderID     string             `bson:"provider_id"`
}

// IsValid checks if instance of User is valid using the validator.
func (u User) IsValid() error {
	validate := validator.New()
	return validate.Struct(u)
}

// UserLoginForm is the input from user on login.
type UserLoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=255"`
}

// IsValid checks if instance of UserLoginForm is valid using the validator.
func (u UserLoginForm) IsValid() error {
	validate := validator.New()
	return validate.Struct(u)
}
