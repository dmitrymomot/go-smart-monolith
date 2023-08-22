package service

import (
	"context"
	"net/http"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/adapters/messagebus"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/adapters/players"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/adapters/storage"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/commands"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/common"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/decorators/events"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/decorators/logger"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/queries"
)

type (
	// Service is a user service facade.
	// It's just a collection of the query and command handlers with the service configuration.
	Service struct {
		GetUser    common.QueryHandler[queries.GetUserQuery, queries.User]
		CreateUser common.CommandHandler[commands.CreateUserCommand]
	}

	// Config holds the user service configuration.
	Config struct {
		Flag              bool
		PlayerSvcEndpoint string
	}

	// low-level abstraction for the storage.
	storageService interface {
		Get(ctx context.Context, key string) (interface{}, error)
		Set(ctx context.Context, key string, value interface{}) error
	}

	// low-level abstraction for the logger.
	loggerX interface {
		Error(err error, kv ...interface{})
	}

	// low-level abstraction for the NATS client.
	natsClient interface {
		Publish(subject string, body []byte) error
	}
)

// NewService returns a new app service instance.
// It's just a factory function that creates a new app service instance.
// It's a good place to apply all the decorators to the app service.
// You can create more different factory functions for different environments
// (e.g. for testing, for production, etc.) with env-specific decorators applied.
func NewService(stor storageService, log loggerX, nc natsClient, cnf Config) Service {
	// Init the user repository.
	userRepo := storage.New(stor)

	// Init the player client.
	playerClient := players.New(players.Config{
		Endpoint: cnf.PlayerSvcEndpoint,
	}, &http.Client{})

	// Init the message bus adapter.
	messageBus := messagebus.NewEventSender(nc)

	// Create the app instance with all the decorators applied.
	userApp := Service{
		GetUser: common.ApplyQueryDecorators(
			queries.GetUser(userRepo, playerClient),
			logger.QueryErrorLogger[queries.GetUserQuery, queries.User](log), // Logs the error if any. So you don't need to care about this in the query handler.
		),
		CreateUser: common.ApplyCommandDecorators(
			commands.CreateUser(userRepo, true),
			logger.CommandErrorLogger[commands.CreateUserCommand](log), // Logs the error if any.
			events.EventSender[commands.CreateUserCommand](messageBus), // Sends the event to the message bus.
		),
	}

	return userApp
}

// httpClient is a low-level abstraction for the HTTP client.
type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

// NewTestService returns a new app service instance for testing.
// It's almost the same as the NewService function but with the test-specific decorators applied.
func NewTestService(stor storageService, log loggerX, nc natsClient, cnf Config, httpc httpClient) Service {
	// Init the user repository.
	userRepo := storage.New(stor)

	// Init the player client.
	playerClient := players.New(players.Config{
		Endpoint: cnf.PlayerSvcEndpoint,
	}, httpc)

	// Init the message bus adapter.
	messageBus := messagebus.NewEventSender(nc)

	// Create the app instance with all the decorators applied.
	userApp := Service{
		GetUser: common.ApplyQueryDecorators(
			queries.GetUser(userRepo, playerClient),
			logger.QueryErrorLogger[queries.GetUserQuery, queries.User](log), // Logs the error if any. So you don't need to care about this in the query handler.
		),
		CreateUser: common.ApplyCommandDecorators(
			commands.CreateUser(userRepo, true),
			logger.CommandErrorLogger[commands.CreateUserCommand](log), // Logs the error if any.
			events.EventSender[commands.CreateUserCommand](messageBus), // Sends the event to the message bus.
		),
	}

	return userApp
}
