package accounts

import (
	"net/http"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	user := auth.AccountOf(r)
	if user == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	followUsername := chi.URLParam(r, "username")
	if followUsername == "" {
		http.Error(w, "username is empty", http.StatusBadRequest)
		return
	}

	followUser, err := h.ar.FindByUsername(ctx, followUsername)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//folowするuserIdとfolowされるuserId
	followerId := user.ID
	followedId := followUser.ID

	if followerId == followedId {
		http.Error(w, "cannot follow yourself", http.StatusBadRequest)
		return
	}

	addFollow := h.rr.AddFollow(ctx, followerId, followedId)
	if err := addFollow; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
