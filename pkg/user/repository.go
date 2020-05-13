package user

import "context"

// Repository handle the CRUD operations with Users.
type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetOne(ctx context.Context, id uint) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id uint, user User) error
	Delete(ctx context.Context, id uint) error
}
