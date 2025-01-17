package usecase

import (
	"code-base-go/internal/domain/entities"
	"code-base-go/internal/repository"
	"errors"
)

type UserUseCaseInterface interface {
	RegisterUser(user *entities.User) error
}

type UserUseCase struct {
	UserRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCaseInterface {
	return &UserUseCase{UserRepo: userRepo}
}

func (u *UserUseCase) RegisterUser(user *entities.User) error {
	if !user.IsValid() {
		return errors.New("invalid user data")
	}

	return u.UserRepo.Save(user)
}