package auth

import (
	"errors"
	"os"
	"time"

	"Logger.Fitness/go-libs/types"
	"github.com/golang-jwt/jwt"
)

// ErrClaimParse is an error of parsing the claim
var ErrClaimParse = errors.New("couldn't parse claim")

// ErrJwtExpired is an error of  JWT expiring
var ErrJwtExpired = errors.New("JWT has expired")

// ErrJwtInvalid is an error of JWT being invalid
var ErrJwtInvalid = errors.New("JWT invalid")

const (
	Issuer             = "Logger.Fitness"
	JWTSecretEnv       = "JWT_SECRET"
	JWTExpirationHours = 24
)

func getJwtSecretFromEnv() string {
	return os.Getenv(JWTSecretEnv)
}

// GenerateJWTFromDbUser function wrapper around GenerateJWTFromClaim.
func GenerateJWTFromDbUser(user types.User) (string, error) {
	claims := &types.JwtClaim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(JWTExpirationHours)).Unix(),
			Issuer:    Issuer,
		},
	}

	signedToken, err := GenerateJWTFromClaim(*claims)
	return signedToken, err
}

// GenerateJWTFromClaim function to generate JWT token from a previous claim.
func GenerateJWTFromClaim(claims types.JwtClaim) (string, error) {
	claims.StandardClaims.ExpiresAt = time.
		Now().
		Local().
		Add(time.Hour * time.Duration(JWTExpirationHours)).
		Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(getJwtSecretFromEnv()))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

// ValidateJWTToken function to validate the JWT token and return the custom claims.
func ValidateJWTToken(userToken string) (claims *types.JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		userToken,
		&types.JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(getJwtSecretFromEnv()), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*types.JwtClaim)
	if !ok {
		err = ErrClaimParse
		return
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		err = ErrJwtExpired
		return
	}
	return
}
