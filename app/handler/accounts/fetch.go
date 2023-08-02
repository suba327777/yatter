package accounts

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *handler) Fetch(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "usernme is empty", http.StatusBadRequest)
	}

	account, err := h.ar.FindByUsername(ctx, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if account == nil {
		http.Error(w, "account is dosen't exist", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
