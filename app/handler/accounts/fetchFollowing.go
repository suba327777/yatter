package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/utils"

	"github.com/go-chi/chi/v5"
)

func (h *handler) FetchFollowing(w http.ResponseWriter, r *http.Request) {

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

	followingAccounts, err := h.rr.FetchFollowingAccounts(ctx, account.ID, q.Limit)
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
