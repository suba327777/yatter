package accounts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const defaultLimit = 40
const maxLimit = 80

type query struct {
	username string `url:"username"`
	limit    int64  `url:"limit,omitempty"`
}

func (h *handler) FetchFollowing(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	q, err := h.QueryValidate(w, r)
	if err != nil {
		http.Error(w, "query error", http.StatusBadRequest)
	}

	account, err := h.ar.FindByUsername(ctx, q.username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followingAccounts, err := h.ar.FetchFollowingAccounts(ctx, account.ID, q.limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(followingAccounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *handler) QueryValidate(w http.ResponseWriter, r *http.Request) (query, error) {
	q := query{}

	queryParams := r.URL.Query()

	q.username = chi.URLParam(r, "username")
	if q.username == "" {
		http.Error(w, "username is empty", http.StatusBadRequest)
	}

	limitStr := queryParams.Get("limit")
	if limitStr == "" {
		q.limit = defaultLimit
	} else {
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			//エラー発生時はデフォルト値をセット
			q.limit = defaultLimit
		} else {
			q.limit = limit
		}
	}

	if q.limit >= maxLimit {
		q.limit = maxLimit
	} else if q.limit <= defaultLimit {
		q.limit = defaultLimit
	}

	return q, nil
}
