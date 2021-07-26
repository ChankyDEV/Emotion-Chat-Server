package services

import (
	"fmt"

	"github.com/ChankyDEV/emotion_server/models"
	"github.com/ChankyDEV/emotion_server/repositories"
)

type IAuthService interface {
	SignUp(email, phone, password string) (*models.User, error)
}

type AuthService struct {
	AuthRepository repositories.IAuthRepository
}

func (as *AuthService) SignUp(email, phone, password string) (*models.User, error) {
	user, err := as.AuthRepository.SignUp(email, phone, password)
	if err != nil {
		return &models.User{}, fmt.Errorf("SOME DATA PROBLEM")
	}
	return &user, nil
}

func NewAuthService(authRepository repositories.IAuthRepository) IAuthService {
	return &AuthService{
		AuthRepository: authRepository,
	}
}
