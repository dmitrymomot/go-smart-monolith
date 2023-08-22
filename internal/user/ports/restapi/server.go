package restapi

import (
	"net/http"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/service"
	"github.com/go-chi/chi/v5"
)

// NewServer creates a new HTTP server.
// It can be used as a standalone server or as a part of a bigger server.
// See cmd/api/main.go for an example.
func NewServer(svc service.Service) http.Handler {
	r := chi.NewRouter()
	// Some more specific middlewares might need to be set on
	// the routes in the user service.
	// Don't place the same middlewares you setup in main() here,
	// because they will be applied to all services and endpoints.
	r.Use(jwtAuthMiddleware)

	// Mount all endpoints here.
	r.Post("/", createUserEndpointHandler(svc))
	r.Get("/{id}", getUserEndpointHandler(svc))

	return r
}

// jwtAuthMiddleware is a middleware that checks if the request is authorized.
// It's just an example of a middleware. Don't pay attention to implementation.
func jwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement JWT auth here.
	})
}
