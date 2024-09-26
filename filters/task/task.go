package task

import (
	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/filters"
	"github.com/cocktail828/go-tools/z/reflectx"
)

type Option func(*Task)

func WithFilters(filters ...filters.Filter) Option {
	return func(t *Task) {
		t.filters = filters
	}
}

type Task struct {
	filters []filters.Filter
}

func New(opts ...Option) (*Task, error) {
	inst := &Task{}
	for _, opt := range opts {
		opt(inst)
	}
	return inst, nil
}

func (t *Task) Run(ctx filters.Context) (err *errcode.Error) {
	for _, f := range t.filters {
		if err := f.Prepare(&ctx); !reflectx.IsNil(err) {
			return err
		}
	}

	for _, f := range t.filters {
		if err := f.Process(&ctx); !reflectx.IsNil(err) {
			return err
		}
	}

	for _, f := range t.filters {
		if err := f.Cleanup(&ctx); !reflectx.IsNil(err) {
			return err
		}
	}

	return nil
}
