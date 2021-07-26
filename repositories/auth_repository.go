package repositories

import "github.com/ChankyDEV/emotion_server/models"

type IAuthRepository interface {
	SignUp(email, phone, password string) (models.User, error)
}

type AuthRepository struct{}

func (*AuthRepository) SignUp(email, phone, password string) (models.User, error) {
	return models.User{
		Uid:      "1",
		Email:    email,
		Phone:    phone,
		Password: password,
	}, nil
}

func NewMongoRepository() IAuthRepository {
	return &AuthRepository{}
}
