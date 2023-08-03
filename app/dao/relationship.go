package dao

import (
	"context"
	"yatter-backend-go/app/domain/object"
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

// FetchFollowingAccounts : フォローしているユーザを取得
func (r *account) FetchFollowingAccounts(ctx context.Context, userId int64, limit int64) ([]*object.Account, error) {

	rows, err := r.db.QueryContext(ctx, "select a.* from `account` as a inner join `relationship` as r on a.id = r.followed_id where r.follower_id = ? limit ?", userId, limit)
	if err != nil {
		return nil, err
	}

	var accounts []*object.Account
	for rows.Next() {
		var account object.Account
		if err := rows.Scan(&account.ID, &account.Username, &account.PasswordHash, &account.DisplayName, &account.Avatar, &account.Header, &account.Note, &account.CreateAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}
