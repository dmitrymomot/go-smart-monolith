package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/queries"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/service"

	"github.com/go-chi/chi/v5"
)

// UserResponse represents the response body for GetUser.
// Note: you can't use the user entity directly as a response,
// follow single responsibility principle and create a separate struct for the response.
type UserResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	PlayerName string `json:"player_name"`
}

// getUserEndpointHandler is a function that handles the HTTP request to get a user.
func getUserEndpointHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request parameters.
		id := chi.URLParam(r, "id")

		// Execute the query.
		user, err := svc.GetUser(r.Context(), queries.GetUserQuery{
			ID: id,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Return the response.
		// Note: you can't pass the user directly to the response,
		// follow single responsibility principle and create a separate struct for the response.
		if err := json.NewEncoder(w).Encode(UserResponse{
			ID:         user.ID,
			Email:      user.Email,
			PlayerName: user.PlayerName,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
