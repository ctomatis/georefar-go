package credentials

import "github.com/golang-jwt/jwt"

type token struct {
	secret, key string
}

func New(key, secret string) *token {
	return &token{
		secret: secret,
		key:    key,
	}
}

func (t *token) jwt() (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{Issuer: t.key}).SignedString([]byte(t.secret))
}

func (t *token) GetAuthentication() (string, error) {
	tk, err := t.jwt()
	return "Bearer " + tk, err
}
