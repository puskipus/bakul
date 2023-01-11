package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("jndi923j9f923k017643vduj")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
