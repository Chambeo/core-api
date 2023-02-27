package service

import (
	"chambeo-core-api/internal/auth/models"
	models2 "chambeo-core-api/internal/models"
)

//go:generate mockgen -source=auth_service.go -destination auth_service_mock.go -package service

type IAuthRepository interface {
	Login(payload models2.LoginPayload) (*string, error)
	SignUp(user *models.User) (*models.User, error)
}

type AuthService struct {
	authRepo IAuthRepository
}

func NewAuthService(authRepo IAuthRepository) AuthService {
	return AuthService{authRepo: authRepo}
}

func (authService *AuthService) Login(payload models2.LoginPayload) (*string, error) {

	token, err := authService.authRepo.Login(payload)
	return token, err

}

func (authService *AuthService) SignUp(userRequest *models.User) (*models.User, error) {

	result, err := authService.authRepo.SignUp(userRequest)

	return result, err
}
