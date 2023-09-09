package models

// AuthenticationRequest represents the JSON structure expected for authentication requests.
type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthenticationResponse represents the JSON structure returned for authentication requests.
type AuthenticationResponse struct {
	ID          int    `json:"account_id,omitempty"`
	Username    string `json:"username,omitempty"`
	AccessToken string `json:"access_token"`
}

// LogoutRequest represents the JSON structure expected for logout requests.
type LogoutRequest struct {
	AccountID string `json:"account_id"`
}
