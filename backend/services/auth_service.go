package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/layunne/todo-list/backend/errors"
	"net/http"
	"time"
)

type AuthService interface {
	Check(token string) bool
	GetId(token string) (string, *errors.Error)
	GetToken(userId string) string
}

func NewAuthService(authSecret string) AuthService {
	return &authService{authSecret: authSecret}
}

type authService struct {
	authSecret string
}

func (s *authService) Check(auth string) bool {

	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error on validating auth token")
		}
		return []byte(s.authSecret), nil
	})
	if err != nil {
		return false
	}

	return token.Valid
}

func (s *authService) GetId(tokenStr string) (string, *errors.Error) {

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id := claims["id"].(string)
		return id, nil
	}
	return "", errors.New(http.StatusUnauthorized, "invalid token")
}

func (s *authService) GetToken(userId string) string {

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := tokenJwt.SignedString([]byte(s.authSecret))

	if err != nil {
		return ""
	}

	return token
}


