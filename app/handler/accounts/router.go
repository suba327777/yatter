package accounts

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
	rr repository.Relationship
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Account, rr repository.Relationship) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar, rr}
	r.Get("/{username}", h.Fetch)
	r.Get("/{username}/following", h.FetchFollowing)
	r.Get("/{username}/followers", h.FetchFollowers)
	r.Post("/", h.Create)
	r.With(auth.Middleware(ar)).Post("/{username}/follow", h.Follow)

	return r
}
