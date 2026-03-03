package di

import (
	"context"
	"database/sql"
)

type UserMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) *UserMySQLRepository {
	return &UserMySQLRepository{db: db}
}

func (r *UserMySQLRepository) Save(ctx context.Context, user *User) error {
	query := `INSERT INTO users (name, email, password) VALUES(?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	return err
}
