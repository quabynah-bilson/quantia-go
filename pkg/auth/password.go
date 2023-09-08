package auth

// PasswordHelper is the interface that wraps the basic password methods.
type PasswordHelper interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
}
