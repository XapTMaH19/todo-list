package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/XapTMaH19/todo-app/internal/models"
	"github.com/XapTMaH19/todo-app/internal/storage"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "asfsfskfdjsa"
	signingKey = "asfasdfsdf"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	storage storage.Authorization
}

func NewAuthService(storage storage.Authorization) *AuthService {
	return &AuthService{storage}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.storage.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	// get user from DB
	user, err := s.storage.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SigningString()
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
