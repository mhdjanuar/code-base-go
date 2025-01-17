package repository

import "code-base-go/internal/domain/entities"

type UserRepository interface {
	GetByID(id int) (*entities.User, error)
	Save(user *entities.User) error
}