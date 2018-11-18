package models

import "github.com/dgrijalva/jwt-go"

type MyCustomClaims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}
