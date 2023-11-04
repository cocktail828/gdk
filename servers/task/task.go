package task

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cocktail828/gdk/v1/message"
	"github.com/cocktail828/gdk/v1/zplugin"
	"github.com/cocktail828/go-tools/z/chain"
	"github.com/cocktail828/go-tools/z/errcode"
	"github.com/cocktail828/go-tools/z/reflectx"
	"golang.org/x/exp/slog"
)

type Option func(*Task)

func WithSrvName(srvname string) Option {
	return func(t *Task) {
		t.srvName = srvname
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(t *Task) {
		t.logger = logger
	}
}

type Task struct {
	srvName   string
	logger    *slog.Logger
	subLogger *slog.Logger
	chain     []zplugin.ZPlugin
	buffer    *bytes.Buffer
}

func New(opts ...Option) (*Task, error) {
	taskInst := &Task{
		buffer: get(),
	}

	for _, opt := range opts {
		opt(taskInst)
	}
	return taskInst, nil
}

func (t *Task) Close() error {
	t.buffer.Reset()
	put(t.buffer)
	return nil
}

func (t *Task) Logger() *slog.Logger {
	if t.subLogger != nil {
		return t.subLogger
	}
	return t.logger
}

func (t *Task) SendBack(d []byte) {
	t.buffer.Write(d)
}

func (t *Task) Run(msg *message.Message) *errcode.Error {
	plugins := []string{}
	zplugin.Traverse(func(res chain.Handler) {
		v := res.(zplugin.ZPlugin)
		if v.Interest(msg) {
			t.chain = append(t.chain, v)
			plugins = append(plugins, v.Name())
		}
	})
	t.logger = t.logger.WithGroup(t.srvName).With("sid", msg.Sid)

	t.logger.Info(fmt.Sprintf("task start with plugins: %v", strings.Join(plugins, ",")))
	err := t.run(msg)
	return err
}

func (t *Task) run(msg *message.Message) *errcode.Error {
	for _, v := range t.chain {
		t.logger.Debug("plugin.Preproc try process request", "plugin", v.Name())
		t.subLogger = t.logger.With("plugin", v.Name(), "fn", "Preproc")
		code, err := v.Preproc(msg, t)
		if code == zplugin.STOP || !reflectx.IsNil(err) {
			return err
		}
	}

	for _, v := range t.chain {
		t.logger.Debug("plugin.Process try process request", "plugin", v.Name())
		t.subLogger = t.logger.With("plugin", v.Name(), "fn", "Process")
		code, err := v.Process(msg, t)
		if code == zplugin.STOP || !reflectx.IsNil(err) {
			return err
		}
	}

	for _, v := range t.chain {
		t.logger.Debug("plugin.Postproc try process request", "plugin", v.Name())
		t.subLogger = t.logger.With("plugin", v.Name(), "fn", "Postproc")
		code, err := v.Postproc(msg, t)
		if code == zplugin.STOP || !reflectx.IsNil(err) {
			return err
		}
	}
	return nil
}
