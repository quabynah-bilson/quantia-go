package account

// Repository is the interface that wraps the basic account methods.
type Repository interface {
	// Register registers a new user.
	Register(username string, password string) error

	// Login logs in a user.
	Login(username string, password string) error
}
