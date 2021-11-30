package types

// LoginResponseDto struct that is getting returned to the user at login
type LoginResponseDto struct {
	Claim JwtClaim `json:"claim"`
	Jwt   string   `json:"jwt"`
}
