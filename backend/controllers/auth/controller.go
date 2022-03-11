//go:generate mockgen -package auth -destination controller_mock.go -source lambda.go
package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"Logger.Fitness/go-libs/auth"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
	models "Logger.Fitness/go-libs/types"
)

const path = "/auth"

// DatabaseInterface database dependency interface
type DatabaseInterface interface {
	AddUser(user models.User) error
	CheckUserBasedOnEmail(lookupEmail string) (bool, error)
	GetUserByEmail(lookupEmail string) (models.User, error)
}

// AuthController struct for auth controller
type AuthController struct {
	database DatabaseInterface
}

// NewAuthController created new auth controller
func NewAuthController(database DatabaseInterface) AuthController {
	return AuthController{
		database: database,
	}
}

// Init the route
func (ctrl AuthController) Init(g *echo.Group) {
	group := g.Group(path)
	group.POST("/register", ctrl.register)
	group.POST("/login", ctrl.login)
	group.POST("/verify_me", ctrl.verifyMe)
}

// Register controller to user registration.
func (ctrl *AuthController) register(c echo.Context) error {
	db := ctrl.database

	// Struct binding
	var newUser types.User
	if bindErr := c.Bind(&newUser); bindErr != nil {
		log.Warn(bindErr)
		return c.JSON(http.StatusBadRequest, res.BadPayload)
	}
	defaultRoles := make(map[string]string)
	defaultRoles["user"] = "user"
	newUser.Roles = defaultRoles
	newUser.Active = false

	// Input validation
	if validationErr := newUser.IsValid(); validationErr != nil {
		log.Info(validationErr)
		return c.JSON(http.StatusBadRequest, validationErr.Error())
	}

	//Check for user
	existingUser, dbErr := db.CheckUserBasedOnEmail(newUser.Email)
	if dbErr != nil {
		log.Warn(dbErr)
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}
	if existingUser {
		return c.String(http.StatusBadRequest, res.UserExists)
	}

	// Hash the password
	hashedPass, hashErr := auth.Hash(newUser.Password)
	if hashErr != nil {
		log.Warn(hashErr)
		return c.NoContent(http.StatusInternalServerError)
	}
	newUser.Password = hashedPass

	// Database insertion
	if dbErr := db.AddUser(newUser); dbErr != nil {
		log.Warn(dbErr)
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}

// Login controller to log in the user and return JWT token.
func (ctrl *AuthController) login(c echo.Context) error {
	db := ctrl.database

	// Struct binding
	var userLogin types.UserLoginForm
	if bindErr := c.Bind(&userLogin); bindErr != nil {
		log.Info(bindErr.Error())
		return c.JSON(http.StatusBadRequest, res.BadPayload)
	}

	// Input validation
	if validationErr := userLogin.IsValid(); validationErr != nil {
		log.Info(validationErr)
		return c.JSON(http.StatusBadRequest, validationErr.Error())
	}

	// Query hashed db pass
	dbUser, dbErr := db.GetUserByEmail(userLogin.Email)
	if dbErr == mongo.ErrNoDocuments {
		return c.NoContent(http.StatusUnauthorized)
	}
	if dbErr != nil {
		log.Warn(dbErr)
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	// Compare hash and plaintext
	if !auth.VerifyHash(dbUser.Password, userLogin.Password) {
		return c.NoContent(http.StatusUnauthorized)
	}

	// Create JWT
	userJwt, jwtErr := auth.GenerateJWTFromDbUser(dbUser)
	if jwtErr != nil {
		log.Info(jwtErr.Error())
		return c.String(http.StatusInternalServerError, res.JwtError)
	}

	// Get user claim and return it
	jwtClaim, err := auth.ValidateJWTToken(userJwt)
	if err != nil {
		log.Info(err.Error())
		return c.String(http.StatusInternalServerError, res.JwtError)
	}

	resp := types.LoginResponseDto{
		Jwt:   userJwt,
		Claim: *jwtClaim,
	}

	return c.JSON(http.StatusOK, resp)
}

// VerifyMe controller to verify the jwt tokens.
func (ctrl *AuthController) verifyMe(c echo.Context) error {
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
	_, jwtErr := auth.ValidateJWTToken(userJwt.Jwt)
	if jwtErr != nil {
		return c.String(http.StatusUnauthorized, res.JwtInvalid)
	}

	return c.NoContent(http.StatusOK)
}
