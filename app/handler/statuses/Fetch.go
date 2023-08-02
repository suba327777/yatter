package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Fetch(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "is is empty", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "id is not number", http.StatusBadRequest)
		return
	}

	status, err := h.sr.FindById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status == nil {
		http.Error(w, "status id is not found", http.StatusNotFound)
		return
	}

	account, err := h.ar.FindById(ctx, status.AccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if account == nil {
		http.Error(w, "account id is not found", http.StatusNotFound)
		return
	}

	req := object.Status{
		ID:       id,
		Account:  *account,
		Content:  status.Content,
		CreateAt: status.CreateAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
