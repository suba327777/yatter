package dao

import (
	"context"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	relationship struct {
		db *sqlx.DB
	}
)

func NewRelationship(db *sqlx.DB) repository.Relationship {
	return &status{db: db}
}

// AddFollow
func (r *status) AddFollow(ctx context.Context, followerId int64, followedId int64) error {
	_, err := r.db.ExecContext(ctx, "insert into `relationship` (`follower_id`, `followed_id`) values (?, ?)", followerId, followedId)
	if err != nil {
		return err
	}

	return nil
}
