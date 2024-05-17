package services

import (
	"context"
	"time"
	"van_thailand_server/config"
	"van_thailand_server/models"
	"van_thailand_server/repositories"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtKey)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func Login(ctx context.Context, userReq models.UserStruct) (string, error) {
	userData, err := repositories.FindUser(ctx, &userReq)
	if err != nil {
		return "", err
	}

	tokenString, err := generateJWT(userData.Username)
	if err != nil {
		// http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return "", err
	}

	if checkPasswordHash(userReq.Password, userData.Password) {
		return tokenString, nil
	}
	return "", nil
}

func Register(ctx context.Context, userReq models.UserStruct) error {
	hashedPassword, err := hashPassword(userReq.Password)
	if err != nil {
		return err
	}
	user := models.UserStruct{
		Username: userReq.Username,
		Password: hashedPassword,
	}
	err = repositories.CreateUser(ctx, &user)
	if err != nil {
		return err
	}
	return err
}
