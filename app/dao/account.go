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
