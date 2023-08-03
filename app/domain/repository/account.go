package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	FindById(ctx context.Context, userId int64) (*object.Account, error)
	Insert(ctx context.Context, username string, passwordHash string) error
}
