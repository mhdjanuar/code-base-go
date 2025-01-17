package repository

import (
	"code-base-go/internal/domain/entities"

	"gorm.io/gorm"
)

type SQLUserRepository struct {
	DB *gorm.DB
}

func NewSQLUserRepo(db *gorm.DB) *SQLUserRepository {
	return &SQLUserRepository{DB: db}
}

func (repo *SQLUserRepository) GetByID(id int) (*entities.User, error) {
	var user entities.User
	if err := repo.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *SQLUserRepository) Save(user *entities.User) error {
	if err := repo.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
