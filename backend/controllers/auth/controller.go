package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"

	"Logger.Fitness/go-libs/auth"
	"Logger.Fitness/go-libs/auth/providers"
	res "Logger.Fitness/go-libs/responses"
	"Logger.Fitness/go-libs/types"
)

const path = "/auth"

// DatabaseInterface database dependency interface
type DatabaseInterface interface {
	AddUser(user types.User) error
	CheckUserBasedOnEmail(lookupEmail string) (bool, error)
	GetUserByEmail(lookupEmail string) (types.User, error)
	GetUserByID(id primitive.ObjectID) (types.User, error)
}

// AuthController struct for auth controller
type AuthController struct {
	database          DatabaseInterface
	googleOauthConfig *oauth2.Config
}

// NewAuthController created new auth controller
func NewAuthController(database DatabaseInterface, googleOauthConfig *oauth2.Config) AuthController {
	return AuthController{
		database:          database,
		googleOauthConfig: googleOauthConfig,
	}
}

// Init the route
func (ctrl AuthController) Init(g *echo.Group) {
	group := g.Group(path)
	group.POST("/register", ctrl.register)
	group.POST("/login", ctrl.login)
	group.POST("/verify_me", ctrl.verifyMe)
	//ctrl.setupOauthProviders(g)
	group.GET("/google/login", ctrl.oauth2Login)
	group.GET("/google/callback", ctrl.oauth2Callback)
}

func (ctrl *AuthController) setupOauthProviders(g *echo.Group) {
	g.GET("/google/login", ctrl.oauth2Login)
	g.GET("/google/callback", ctrl.oauth2Callback)
}

// register controller to user registration.
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

	// Input validation
	if inputErr := newUser.IsValid(); inputErr != nil {
		return inputErr
	}

	// Database insertion
	if dbErr := db.AddUser(newUser); dbErr != nil {
		log.Warn(dbErr)
		return c.String(http.StatusInternalServerError, res.DatabseError)
	}

	return c.NoContent(http.StatusOK)
}

// login controller to log in the user and return JWT token.
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

// verifyMe controller to verify the jwt tokens.
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

// TODO: need to fiund a way to decouple the oauth2 functions from the specific providers
func (ctrl *AuthController) oauth2Login(c echo.Context) error {
	// TODO: look up what state is meant to be and change it
	// NOTE: this needs to be a random string, it protects against CSRF attacks, but need to read more about it
	//	see https://datatracker.ietf.org/doc/html/rfc6749#section-10.12
	loginURL := ctrl.googleOauthConfig.AuthCodeURL("state")
	return c.String(http.StatusOK, loginURL)
}

func (ctrl *AuthController) oauth2Callback(c echo.Context) error {
	db := ctrl.database
	state := c.FormValue("state")
	code := c.FormValue("code")

	log.Printf("state= %s", state)
	log.Printf("code= %s", code)

	googleUser, err := providers.GetUserFromGoogle(state, code, ctrl.googleOauthConfig)
	if err != nil {
		log.Printf("Error getting user info from provider: %v", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	_, err = db.GetUserByID(googleUser.ID)
	if err != nil && err != mongo.ErrNoDocuments {

		log.Info(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	if err == mongo.ErrNoDocuments {
		// TODO: need an "add or update" function
		err = db.AddUser(googleUser)
		if err != nil {

			log.Info(err.Error())
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	// Generating JWT
	userJwt, err := auth.GenerateJWTFromDbUser(googleUser)
	if err != nil {

		log.Info(err.Error())
		return c.NoContent(http.StatusInternalServerError)
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
