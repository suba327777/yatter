package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Relationship interface {
	AddFollow(ctx context.Context, followerId int64, followedId int64) error
	FetchFollowing(ctx context.Context, userId int64, limit int64) ([]*object.Account, error)
	FetchFollowers(ctx context.Context, userId int64, maxId int64, sinceId int64, limit int64) ([]*object.Account, error)
}
