package types

// User is a direct representation of the mongo document.
type User struct {
	ID             string            `json:"id"`
	Email          string            `json:"email"`
	Username       string            `json:"username"`
	Password       string            `json:"password"`
	Active         bool              `json:"active"`
	Roles          map[string]string `json:"roles"`
	ProfilePicture string            `json:"profile_picture"`
	Provider       string            `json:"provider"`
	ProviderID     string            `json:"provider_id"`
}

// UserLoginForm is the input from user on login.
type UserLoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
