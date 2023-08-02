package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	FindById(ctx context.Context, id int64) (*object.Status, error)
	Insert(ctx context.Context, userId int64, content string) (int64, error)
}
