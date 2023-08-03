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

const defaultLimit = 40
const maxLimit = 80

func ValidateQuery(r *http.Request) (query, error) {
	q := query{}

	queryParams := r.URL.Query()

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
