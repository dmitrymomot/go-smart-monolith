package common_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/common"

	"github.com/stretchr/testify/require"
)

func TestApplyCommandDecorators(t *testing.T) {
	t.Parallel()

	// Define a command type.
	type TestCommand struct{ Val string }
	// Define an event type.
	type TestEvent struct{ Val string }

	// Create a command handler function.
	handler := func(ctx context.Context, cmd TestCommand) ([]interface{}, error) {
		if cmd.Val != "decorated:test" {
			return nil, fmt.Errorf("unexpected command id: want=%v, got=%v", "decorated:test", cmd.Val)
		}
		return []interface{}{TestEvent(cmd)}, nil
	}

	// Create a command decorator function.
	decorator := func(handler common.CommandHandler[TestCommand]) common.CommandHandler[TestCommand] {
		return func(ctx context.Context, cmd TestCommand) ([]interface{}, error) {
			cmd.Val = "decorated:" + cmd.Val
			return handler(ctx, cmd)
		}
	}

	// Apply the decorator to the handler.
	handler = common.ApplyCommandDecorators(handler, decorator)

	// Call the handler.
	e, err := handler(context.Background(), TestCommand{Val: "test"})
	require.NoError(t, err)
	require.Len(t, e, 1)
	require.Equal(t, TestEvent{Val: "decorated:test"}, e[0])
}
