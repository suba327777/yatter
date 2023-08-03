package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

// FindById : IDからユーザを取得
func (r *account) FindById(ctx context.Context, userId int64) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from `account` where `id` = ?", userId).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}
	return entity, nil
}

// Insert : ユーザを追加
func (r *account) Insert(ctx context.Context, username string, passwordHash string) error {
	_, err := r.db.ExecContext(ctx, "insert into `account` (`username`, `password_hash`) values (?, ?)", username, passwordHash)
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
