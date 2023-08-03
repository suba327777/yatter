package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Relationship interface {
	AddFollow(ctx context.Context, followerId int64, followedId int64) error
	FetchFollowingAccounts(ctx context.Context, userId int64, limit int64) ([]*object.Account, error)
}
