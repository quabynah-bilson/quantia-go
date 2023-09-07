package adapters

// AppDatabase is the interface that wraps the basic database operations.
// It has a generic interface `t` that represents the type of the data that will be stored.
type AppDatabase[t interface{}] interface {

	// Get retrieves a record from the database.
	Get(key string) (*t, error)

	// Set saves a record into the database.
	Set(key string, value *t) error

	// Delete removes a record from the database.
	Delete(key string) error
}
