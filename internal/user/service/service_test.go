package service_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/commands"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/queries"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/domain"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// storageService is a mock of the storageService interface.
type storageService struct {
	mock.Mock
}

// Get is a mock implementation of the Get method.
func (m *storageService) Get(ctx context.Context, key string) (interface{}, error) {
	args := m.Called(ctx, key)
	return args.Get(0), args.Error(1)
}

// Set is a mock implementation of the Set method.
func (m *storageService) Set(ctx context.Context, key string, value interface{}) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

// loggerX is a mock of the loggerX interface.
type loggerX struct {
	mock.Mock
}

// Error is a mock implementation of the Error method.
func (m *loggerX) Error(err error, kv ...interface{}) {
	m.Called(err, kv)
}

// natsClient is a mock of the natsClient interface.
type natsClient struct {
	mock.Mock
}

// Publish is a mock implementation of the Publish method.
func (m *natsClient) Publish(subject string, body []byte) error {
	args := m.Called(subject, body)
	return args.Error(0)
}

// httpClient is a mock of the httpClient interface.
type httpClient struct {
	mock.Mock
}

// Get is a mock implementation of the Get method.
func (m *httpClient) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestService_CreateUser(t *testing.T) {
	// Test data.
	email := "test@mail.dev"

	// Create a new mock for the storageService.
	stor := new(storageService)
	stor.On("Get", mock.Anything, email).Return(nil, errors.New("error"))
	stor.On("Set", mock.Anything, email, mock.Anything).Return(nil)

	// Set mocks.
	log := new(loggerX)
	httpc := new(httpClient)

	// Create a new mock for the natsClient.
	nc := new(natsClient)
	nc.On("Publish", mock.Anything, mock.Anything).Return(nil)

	// Create a new service instance.
	svc := service.NewTestService(stor, log, nc, service.Config{}, httpc)

	// Call the CreateUser command handler.
	events, err := svc.CreateUser(context.Background(), commands.CreateUserCommand{
		Email:    email,
		Password: "password",
	})

	// Assert that the CreateUser command handler returns no error.
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.IsType(t, commands.UserCreatedEvent{}, events[0])

	// Assert that mocks expectations were met.
	httpc.AssertExpectations(t)
	stor.AssertExpectations(t)
	log.AssertExpectations(t)
	nc.AssertExpectations(t)
}

func TestService_GetUser(t *testing.T) {
	// Test data.
	email := "test@mail.dev"
	user := domain.NewUser(email, "password")

	// Create a new mock for the storageService.
	stor := new(storageService)
	stor.On("Get", mock.Anything, user.ID).Return(user, nil)

	// Create a new mock for the loggerX.
	log := new(loggerX)
	// Create a new mock for the natsClient.
	nc := new(natsClient)

	// Create a new mock for the httpClient.
	httpc := new(httpClient)
	httpc.On("Get", mock.AnythingOfType("string")).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"user_id":"user_id","player_name":"player_name"}`))),
	}, nil)

	// Create a new service instance.
	svc := service.NewTestService(stor, log, nc, service.Config{
		Flag:              true,
		PlayerSvcEndpoint: "http://localhost:8080",
	}, httpc)

	// Call the CreateUser command handler.
	resp, err := svc.GetUser(context.Background(), queries.GetUserQuery{
		ID: user.ID,
	})

	// Assert that the CreateUser command handler returns no error.
	assert.NoError(t, err)
	assert.Equal(t, user.ID, resp.ID)
	assert.Equal(t, "player_name", resp.PlayerName)
	assert.Equal(t, email, resp.Email)

	// Assert that mocks expectations were met.
	httpc.AssertExpectations(t)
	stor.AssertExpectations(t)
	log.AssertExpectations(t)
	nc.AssertExpectations(t)
}
