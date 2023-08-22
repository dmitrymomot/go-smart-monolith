package common

import "context"

// CommandHandler is a command function that can be executed by the service.
type CommandHandler[Cmd any] func(ctx context.Context, cmd Cmd) ([]interface{}, error)

// CommandDecorator is a function that wraps a command handler.
type CommandDecorator[Cmd any] func(CommandHandler[Cmd]) CommandHandler[Cmd]

// ApplyCommandDecorators applies the given command decorators to the given command handler.
func ApplyCommandDecorators[Cmd any](
	handler CommandHandler[Cmd],
	decorators ...CommandDecorator[Cmd],
) CommandHandler[Cmd] {
	for _, decorator := range decorators {
		handler = decorator(handler)
	}
	return handler
}
