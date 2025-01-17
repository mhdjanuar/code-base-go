package repository

import (
	"code-base-go/internal/domain/entities"
	"database/sql"
)

type SQLUserRepository struct {
	DB *sql.DB
}

func NewSQLUserRepo(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{DB: db}
}

func (r *SQLUserRepository) GetByID(id int) (*entities.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = $1"
	row := r.DB.QueryRow(query, id)

	user := &entities.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *SQLUserRepository) Save(user *entities.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}