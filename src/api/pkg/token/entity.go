package token

// Session represents a user session.
type Session struct {
	ID        string `json:"id" bson:"_id"`
	AccountID string `json:"account_id" bson:"account_id"`
	Token     string `json:"token" bson:"token"`
}
