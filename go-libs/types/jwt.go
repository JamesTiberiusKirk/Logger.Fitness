package types

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JwtClaim adds ID, username, email and roles as a claim to the token.
type JwtClaim struct {
	ID       primitive.ObjectID
	Username string            `json:"username"`
	Email    string            `json:"email"`
	Roles    map[string]string `json:"roles"`
	jwt.StandardClaims
}

// JwtDto DTO for the jwt alone
type JwtDto struct {
	Jwt string `json:"jwt" validate:"required"`
}

// IsValid validates an instance of JwtDto using the validator
func (j JwtDto) IsValid() error {
	validate := validator.New()
	return validate.Struct(j)
}
