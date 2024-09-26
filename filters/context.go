package filters

import (
	"context"
	"log/slog"
	"math"
)

type Context struct {
	context context.Context
	Type    string
	index   int16
	logger  *slog.Logger
}

func (ctx *Context) Context() context.Context {
	return ctx.context
}

func (ctx *Context) Abort() {
	ctx.index = math.MaxInt16
}

func (ctx *Context) Logger() *slog.Logger {
	return ctx.logger
}
