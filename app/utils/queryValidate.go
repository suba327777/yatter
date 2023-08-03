package utils

import (
	"net/http"
	"strconv"
)

type query struct {
	MaxId   int64 `url:"max_id,omitempty"`
	SinceId int64 `url:"since_id,omitempty"`
	Limit   int64 `url:"limit,omitempty"`
}

const defaultMaxId = 100
const defaultSinceId = 0
const defaultLimit = 40
const maxLimit = 80

func ValidateQuery(r *http.Request) (query, error) {
	q := query{}

	queryParams := r.URL.Query()

	maxIdStr := queryParams.Get("max_id")
	if maxIdStr == "" {
		// 空の場合はdefault値をセット
		q.MaxId = defaultMaxId
	} else {
		maxId, err := strconv.ParseInt(maxIdStr, 10, 64)
		if err != nil {
			return q, err
		} else {
			q.MaxId = maxId
		}
	}

	sinceIdStr := queryParams.Get("since_id")
	if sinceIdStr == "" {
		q.SinceId = defaultSinceId
	} else {
		sinceId, err := strconv.ParseInt(sinceIdStr, 10, 64)
		if err != nil {
			return q, err
		} else {
			q.SinceId = sinceId
		}
	}

	limitStr := queryParams.Get("limit")
	if limitStr == "" {
		q.Limit = defaultLimit
	} else {
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			//エラー発生時はデフォルト値をセット
			q.Limit = defaultLimit
		} else {
			q.Limit = limit
		}
	}

	if q.Limit >= maxLimit {
		q.Limit = maxLimit
	} else if q.Limit <= defaultLimit {
		q.Limit = defaultLimit
	}

	return q, nil

}
