package repository

import (
	"fitness/internal/model"

	"github.com/jmoiron/sqlx"

)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
}

func (r *userRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE id=$1`
	err := r.db.Get(&user, query, id)
	return &user, err
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE email=$1`
	err := r.db.Get(&user, query, email)
	return &user, err
}

func (r *userRepository) Update(user *model.User) error {
	query := `UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	return err
}

func (r *userRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}
