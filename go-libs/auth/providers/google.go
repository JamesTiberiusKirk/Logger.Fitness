package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"Logger.Fitness/go-libs/types"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleContents struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	IsEmailVeryfied bool   `json:"verified_email"`
	Name            string `json:"name"`
	GivenName       string `json:"given_name"`
	FamilyName      string `json:"family_name"`
	ProfilePicture  string `json:"picture"`
	Locale          string `json:"locale"`
}

// TODO: create a struct for this to have it mockable on the controller side?
// TODO: have vars for this config coming from the env
// SetupGogoleConfig - ...
func SetupGogoleConfig(clientID, clientSecret string) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  "http://dev.logger.fitness:8100/oauth/redirect",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

// GetUserFromGoogle - gets user info from google api and returns types.User
func GetUserFromGoogle(state string, code string, googleOauthConfig *oauth2.Config) (types.User, error) {
	googleUser := types.User{}

	if state != "state" {
		return googleUser, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return googleUser, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return googleUser, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return googleUser, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	var googleUserContent googleContents
	err = json.Unmarshal(contents, &googleUserContent)
	if err != nil {
		return googleUser, err
	}

	defaultRoles := make(map[string]string)
	defaultRoles["user"] = "user"

	googleUser = types.User{
		Email:          googleUserContent.Email,
		Username:       googleUserContent.Name,
		ProfilePicture: googleUserContent.ProfilePicture,
		Provider:       "google",
		ProviderID:     googleUserContent.ID,
		Roles:          defaultRoles,
	}

	return googleUser, nil
}
