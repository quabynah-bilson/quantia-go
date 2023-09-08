package auth

// Account represents a user account.
type Account struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// Session represents a user session.
type Session struct {
	ID        string `json:"id" bson:"_id"`
	AccountID string `json:"account_id" bson:"account_id"`
	Token     string `json:"token" bson:"token"`
}
