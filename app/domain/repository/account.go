package repository

import (
	"context"
	"time"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// TODO: Add Other APIs
	Insert(ctx context.Context, username string, passwordHash string, createAt time.Time) error
}
