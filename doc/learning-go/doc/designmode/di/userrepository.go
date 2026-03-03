package di

import (
	"context"
)

type User struct {
	ID       int    `json:"id" `
	Name     string `json:"name" `
	Email    string `json:"email"`
	Password string `json:"password" `
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
}
