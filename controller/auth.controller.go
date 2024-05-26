package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"van_thailand_server/configdata"
	"van_thailand_server/models"
	"van_thailand_server/services"

	"github.com/dgrijalva/jwt-go"
)

func HandleAuth(ctx context.Context) {
	http.Handle("/admin", authMiddleware())
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var userReq *models.UserStruct
		if r.Method == http.MethodPost {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Invalid body: ", err)
				return
			}
			err = json.Unmarshal(body, &userReq)
			if err != nil {
				log.Println("Unmarshal failed: ", err)
				return
			}
			if userReq != nil {
				token, err := services.Login(ctx, *userReq)
				if err != nil {
					log.Println("Login failed: ", err)
					return
				}
				http.SetCookie(w, &http.Cookie{
					Name:     "session_token",
					Value:    token,
					Path:     "/",
					HttpOnly: true,
					Secure:   true,
				})
				http.Redirect(w, r, "/admin", http.StatusFound)
				return
			} else {
				log.Println("No body given.")
				ReturnFailed(w)
				return
			}
		} else {
			log.Println(w, "Method not allowed.")
			return
		}
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var userReq *models.UserStruct
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Invalid body: ", err)
				return
			}
			err = json.Unmarshal(body, &userReq)
			if err != nil {
				log.Println("Unmarshal failed: ", err)
				return
			}
			if userReq != nil {
				err = services.Register(ctx, *userReq)
				if err != nil {
					log.Println("Register failed: ", err)
					return
				}
			} else {
				log.Println("No body given.")
				ReturnFailed(w)
				return
			}
		}
	})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.SetCookie(w, &http.Cookie{
				Name:   "session_token",
				Value:  "",
				Path:   "/",
				MaxAge: -1,
			})
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	})
}

func authMiddleware() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if r.URL.Path != "/login" {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			return
		}
		err = verifyJWT(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	})
}

func verifyJWT(tokenString string) error {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return configdata.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return fmt.Errorf("invalid token signature")
		}
		return fmt.Errorf("invalid token")
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
