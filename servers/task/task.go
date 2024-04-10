package task

import (
	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/zplugin"
	"github.com/cocktail828/gdk/v1/zplugin/consts"
)

type Option func(*Task)

func WithPlugins(names ...string) Option {
	return func(t *Task) {
		t.names = names
	}
}

type Task struct {
	names []string
}

func New(opts ...Option) (*Task, error) {
	inst := &Task{}
	for _, opt := range opts {
		opt(inst)
	}
	return inst, nil
}

func (t *Task) Run(ctx zplugin.Context) (err *errcode.Error) {
	var code consts.State
	zplugin.Traverse(func(p zplugin.ZPlugin) *errcode.Error {
		if code, err = p.Prepare(&ctx); code == consts.STOP || err != nil {
			return err
		}
		return nil
	}, t.names...)

	zplugin.Traverse(func(p zplugin.ZPlugin) *errcode.Error {
		if code, err = p.Process(&ctx); code == consts.STOP || err != nil {
			return err
		}
		return nil
	}, t.names...)

	zplugin.Traverse(func(p zplugin.ZPlugin) *errcode.Error {
		if code, err = p.Cleanup(&ctx); code == consts.STOP || err != nil {
			return err
		}
		return nil
	}, t.names...)

	return nil
}
