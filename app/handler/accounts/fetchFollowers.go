package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/utils"

	"github.com/go-chi/chi/v5"
)

func (h *handler) FetchFollowers(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "username is empty", http.StatusBadRequest)
		return
	}

	q, err := utils.ValidateQuery(r)
	if err != nil {
		http.Error(w, "query error", http.StatusBadRequest)
	}

	account, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followersAccounts, err := h.rr.FetchFollowers(ctx, account.ID, q.MaxId, q.SinceId, q.Limit)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(followersAccounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
