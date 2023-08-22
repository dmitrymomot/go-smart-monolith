package queries_test

import (
	"context"
	"testing"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/queries"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/domain"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// mockGetUserRepository is a mock of the getUserRepository interface.
type mockGetUserRepository struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: email
func (m *mockGetUserRepository) GetUserByID(ctx context.Context, email string) (domain.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(domain.User), args.Error(1)
}

// mockPlayersSvcClient is a mock of the playersSvcClient interface.
type mockPlayersSvcClient struct {
	mock.Mock
}

// GetPlayer provides a mock function with given fields: key
func (m *mockPlayersSvcClient) GetPlayer(ctx context.Context, key string) (domain.Player, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(domain.Player), args.Error(1)
}

func TestGetUser_Success(t *testing.T) {
	// Prepare test data.
	email := "test@mail.dev"
	pname := "player name"
	user := domain.NewUser(email, "password")
	player := domain.NewPlayer(user.ID, pname)

	// Setup mocks.
	repo := new(mockGetUserRepository)
	repo.On("GetUserByID", mock.Anything, email).Return(user, nil)

	playersClient := new(mockPlayersSvcClient)
	playersClient.On("GetPlayer", mock.Anything, user.ID).Return(player, nil)

	// Create the handler.
	handler := queries.GetUser(repo, playersClient)

	// Setup query.
	query := queries.GetUserQuery{
		ID: email,
	}

	// Call the method under test.
	res, err := handler(context.Background(), query)
	require.NoError(t, err)
	require.Equal(t, queries.User{
		ID:         user.ID,
		Email:      email,
		PlayerName: pname,
	}, res)

	// Verify mocks.
	repo.AssertExpectations(t)
	playersClient.AssertExpectations(t)
}
