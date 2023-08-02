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
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// FindById : Idからstatusを取得
func (r *status) FindById(ctx context.Context, userId int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from `status` where `id` = ?", userId).StructScan(entity)
	if err != nil {
		//結果セットが空の場合
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil

}

// Insert : statusを追加
func (r *status) Insert(ctx context.Context, userId int64, content string) (int64, error) {
	ids, err := r.db.ExecContext(ctx, "insert into `status` (`account_id`, `content`) values (?,?)", userId, content)
	if err != nil {
		return 0, err
	}

	//複数の値を返さないように単一の値を返す
	id, err := ids.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
