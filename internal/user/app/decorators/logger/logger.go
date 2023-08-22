package logger

import (
	"context"
	"fmt"
	"strings"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/common"
)

// Logger is a function that logs a message.
type Logger interface {
	Error(err error, kv ...interface{})
}

// QueryErrorLogger is a decorator that logs query errors.
func QueryErrorLogger[Qry any, Rsp any](logger Logger) common.QueryDecorator[Qry, Rsp] {
	return func(next common.QueryHandler[Qry, Rsp]) common.QueryHandler[Qry, Rsp] {
		return func(ctx context.Context, qry Qry) (Rsp, error) {
			rsp, err := next(ctx, qry)
			if err != nil {
				logger.Error(err, "query", fullyQualifiedStructName(qry))
			}
			return rsp, err
		}
	}
}

// CommandErrorLogger is a decorator that logs command errors.
func CommandErrorLogger[Cmd any](logger Logger) common.CommandDecorator[Cmd] {
	return func(next common.CommandHandler[Cmd]) common.CommandHandler[Cmd] {
		return func(ctx context.Context, cmd Cmd) ([]interface{}, error) {
			e, err := next(ctx, cmd)
			if err != nil {
				logger.Error(err, "command", fullyQualifiedStructName(cmd))
			}
			return e, err
		}
	}
}

// fullyQualifiedStructName name returns object name in format [package].[type name].
// It ignores if the value is a pointer or not.
func fullyQualifiedStructName(v interface{}) string {
	s := fmt.Sprintf("%T", v)
	s = strings.TrimLeft(s, "*")

	return s
}
