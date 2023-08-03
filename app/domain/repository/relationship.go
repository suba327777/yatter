package repository

import "context"

type Relationship interface {
	AddFollow(ctx context.Context, followerId int64, followedId int64) error
}
