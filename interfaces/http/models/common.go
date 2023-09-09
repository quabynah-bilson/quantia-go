package models

// APIResponse represents the format of responses sent by the API.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

// APIError represents an error in the API response.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
