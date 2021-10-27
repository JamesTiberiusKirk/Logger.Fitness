package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"Logger.Fitness/backend/db"
	lib "Logger.Fitness/go-libs"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
)

// Register controller to user registration.
func Register(c echo.Context) error {
	db := c.Get("db").(*db.Client)

	// Struct binding
	var nUser types.User
	if bindErr := c.Bind(&nUser); bindErr != nil {
		log.Info(bindErr)
		return c.JSON(http.StatusBadRequest, res.BAD_PAYLOAD)
	}
	defaultRoles := make([]string, 1)
	defaultRoles[0] = "user"
	nUser.Roles = defaultRoles

	// Input validatin
	if validationErr := nUser.IsValid(); validationErr != nil {
		log.Info(validationErr)
		return c.JSON(http.StatusBadRequest, validationErr.Error())
	}

	//Check for user
	existingUser, dbErr := db.CheckUser(nUser.Email)
	if dbErr != nil {
		log.Fatal(dbErr)
		return c.String(http.StatusInternalServerError, res.DATABASE_ERR)
	}
	if existingUser {
		return c.String(http.StatusBadRequest, res.USER_EXISTS)
	}

	// Hash the password
	hashedPass, hashErr := lib.Hash(nUser.Password)
	if hashErr != nil {
		log.Fatal(hashErr)
		return c.String(http.StatusInternalServerError, res.INTERNAL_SERVER_ERR)
	}
	nUser.Password = hashedPass

	// Database insertion
	if dbErr := db.AddUser(nUser); dbErr != nil {
		log.Fatal(dbErr)
		return c.String(http.StatusInternalServerError, res.DATABASE_ERR)
	}

	return c.String(http.StatusOK, res.SUCCESSFULLY_REGISTERED)
}

// Login controller to log in the user and return JWT token.
func Login(c echo.Context) error {
	db := c.Get("db").(*db.Client)

	// Struct binding
	var userLogin types.UserLoginForm
	if bindErr := c.Bind(&userLogin); bindErr != nil {
		log.Info(bindErr.Error())
		return c.JSON(http.StatusBadRequest, res.BAD_PAYLOAD)
	}

	// Input validation
	if validationErr := userLogin.IsValid(); validationErr != nil {
		log.Info(validationErr)
		return c.JSON(http.StatusBadRequest, validationErr.Error())
	}

	// Query hashed db pass
	dbUser, dbErr := db.GetUserByEmail(userLogin.Email)
	if dbErr == mongo.ErrNoDocuments {
		return c.String(http.StatusUnauthorized, res.UNAUTHORIZED)
	}
	if dbErr != nil {
		log.Fatal(dbErr)
		return c.String(http.StatusInternalServerError, res.DATABASE_ERR)
	}

	// Compare hash and plaintext
	if !lib.VerifyHash(dbUser.Password, userLogin.Password) {
		return c.String(http.StatusUnauthorized, res.UNAUTHORIZED)
	}

	// Create JWT
	userJwt, jwtErr := lib.GenerateJWT(dbUser)
	if jwtErr != nil {
		log.Info(jwtErr.Error())
		return c.String(http.StatusInternalServerError, res.JWT_ERR)
	}

	resp := types.LoginResponseDto{
		Jwt:     userJwt,
		Message: res.LOGGED_IN,
	}

	return c.JSON(http.StatusOK, resp)
}

// VerifyMe controller to verify the jwt tokens.
func VerifyMe(c echo.Context) error {
	var userJwt types.JwtDto
	if bindErr := c.Bind(&userJwt); bindErr != nil {
		log.Info(bindErr)
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	// Input validation
	if validateErr := userJwt.IsValid(); validateErr != nil {
		log.Info(validateErr)
		return c.JSON(http.StatusBadRequest, validateErr.Error())
	}

	// Validate JTW token
	_, jwtErr := lib.ValidateJwtToken(userJwt.Jwt)
	if jwtErr != nil {
		return c.String(http.StatusUnauthorized, res.JWT_INVALID)
	}

	return c.String(http.StatusOK, res.OK)
}
