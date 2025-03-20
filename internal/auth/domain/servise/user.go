package service

import (
	"auth/domain/entity"
	"context"
)

// UserStorage defines the interface for user storage operations
type UserStorage interface {
	Create(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, email string) error
}

type RegistrationService struct {
	storage UserStorage
}

func (r *RegistrationService) NewRegistrationService(storage UserStorage) *RegistrationService {
	return &RegistrationService{storage: storage}
}

func (r *RegistrationService) CreateUser(ctx context.Context, user *entity.User) error {
	err := r.storage.Create(ctx, user)
	return err
}
