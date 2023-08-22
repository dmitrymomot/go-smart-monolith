package common_test

import (
	"context"
	"testing"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/common"
)

// Test query decorators applied to a query handler function.
func TestApplyQueryDecorators(t *testing.T) {
	t.Parallel()

	// Define a query type.
	type TestQuery struct{ Val string }

	// Create a query handler function.
	handler := func(ctx context.Context, qry TestQuery) (string, error) {
		return qry.Val, nil
	}

	// Create a query decorator function.
	decorator := func(handler common.QueryHandler[TestQuery, string]) common.QueryHandler[TestQuery, string] {
		return func(ctx context.Context, qry TestQuery) (string, error) {
			resp, err := handler(ctx, qry)
			if err == nil {
				resp = "decorated:" + resp
			}
			return resp, err
		}
	}

	// Apply the decorator to the handler.
	handler = common.ApplyQueryDecorators(handler, decorator)

	// Call the handler.
	rsp, err := handler(context.Background(), TestQuery{Val: "test"})
	// Check the result.
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if rsp != "decorated:test" {
		t.Errorf("unexpected response: want=%v, got=%v", "decorated:test", rsp)
	}
}
