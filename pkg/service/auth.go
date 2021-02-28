package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Dimadetected/clientsCRUD"
	"github.com/Dimadetected/clientsCRUD/pkg/repository"
)

const (
	salt = "sdfsfcbns;dher123ndvsn;if12"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user clientsCRUD.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
func (s *AuthService) ParseToken(token string) (int, error) {
	return 0, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
