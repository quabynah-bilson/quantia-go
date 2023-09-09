package token

import (
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"github.com/quabynah-bilson/quantia/pkg/token"
	"os"
	"time"
)

var (
	tokenIssuer   = "Quantia Bank"
	tokenSubject  = "quantia_auth_token"
	tokenAudience = "quantia_audience"
)

// PasetoTokenizerHelper is the token helper implementation for Paseto.
type PasetoTokenizerHelper struct {
	token.TokenizerHelper
}

// NewPasetoTokenizerHelper creates a new PasetoTokenizerHelper
func NewPasetoTokenizerHelper() token.TokenizerHelper {
	return &PasetoTokenizerHelper{}
}

// GenerateToken generates a token for the given claim.
func (p *PasetoTokenizerHelper) GenerateToken(claim string) (string, error) {
	// generate a new symmetric key
	now := time.Now()
	jsonToken := paseto.JSONToken{
		Audience:   tokenAudience,
		Issuer:     tokenIssuer,
		Jti:        uuid.NewString(), // unique identifier for the token
		Subject:    tokenSubject,
		IssuedAt:   now,
		Expiration: now.Add(1 * time.Hour), // const time for banking apps (1 hour)
		NotBefore:  now,
	}

	// add custom claim to the token
	jsonToken.Set("claim", claim)

	// encrypt data
	return paseto.NewV2().Encrypt([]byte(os.Getenv("PASETO_SECRET")), jsonToken, os.Getenv("PASETO_FOOTER"))
}

// ValidateToken validates the given token.
func (p *PasetoTokenizerHelper) ValidateToken(rawToken string) error {
	// decrypt token
	var newJSONToken paseto.JSONToken
	var newFooter string
	if err := paseto.NewV2().Decrypt(rawToken, []byte(os.Getenv("PASETO_SECRET")), &newJSONToken, &newFooter); err != nil {
		return token.ErrInvalidToken
	}

	// validate token
	if err := newJSONToken.Validate(paseto.ValidAt(time.Now())); err != nil {
		return token.ErrTokenExpired
	}
	if err := newJSONToken.Validate(paseto.IssuedBy(tokenIssuer)); err != nil {
		return token.ErrInvalidClaim
	}
	if err := newJSONToken.Validate(paseto.ForAudience(tokenAudience)); err != nil {
		return token.ErrInvalidClaim
	}
	if err := newJSONToken.Validate(paseto.Subject(tokenSubject)); err != nil {
		return token.ErrInvalidClaim
	}

	return nil
}
