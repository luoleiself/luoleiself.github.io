package di

import "context"

type UserCase struct {
	UserRepository UserRepository
}

func NewUserCase(ur UserRepository) *UserCase {
	return &UserCase{UserRepository: ur}
}
func (u *UserCase) Create(ctx context.Context, user *User) error {
	return u.UserRepository.Save(ctx, user)
}
