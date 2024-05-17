package models

import "github.com/dgrijalva/jwt-go"

type UserStruct struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
