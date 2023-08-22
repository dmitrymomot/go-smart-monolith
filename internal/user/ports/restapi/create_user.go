package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/commands"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/service"
)

// createUserEndpointHandler is a function that handles the HTTP request to create a user.
func createUserEndpointHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body.
		payload := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: Validate the request payload....

		// Execute the command.
		if _, err := svc.CreateUser(r.Context(), commands.CreateUserCommand{
			Email:    payload.Email,
			Password: payload.Password,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return 201 Created.
		w.WriteHeader(http.StatusCreated)
	}
}
