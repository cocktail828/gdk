package task

import (
	"bytes"
	"strings"

	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/message"
	"github.com/cocktail828/gdk/v1/zplugin"
	"github.com/cocktail828/go-tools/z/reflectx"
	"github.com/pkg/errors"
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

func WithPlugins(ps []zplugin.ZPlugin) Option {
	return func(t *Task) {
		t.plugins = ps
	}
}

type Task struct {
	// inherit variables
	srvName string
	logger  *slog.Logger
	plugins []zplugin.ZPlugin
	// local variables
	subLogger *slog.Logger
	buffer    *bytes.Buffer
}

func New(opts ...Option) (*Task, error) {
	inst := &Task{
		buffer: get(),
	}
	for _, opt := range opts {
		opt(inst)
	}
	if inst.logger == nil {
		return nil, errors.Errorf("invalid task for missing logger")
	}
	if inst.srvName == "" {
		return nil, errors.Errorf("invalid task for missing srvName")
	}
	if len(inst.plugins) == 0 {
		return nil, errors.Errorf("invalid task for empty plugins")
	}
	return inst, nil
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
	for _, p := range t.plugins {
		plugins = append(plugins, p.Name())
	}
	t.logger = t.logger.WithGroup(t.srvName).With("sid", msg.Sid)
	t.logger.Info("task.Run", "op", "task start with plugins", "plugins", strings.Join(plugins, ","))
	return t.run(msg)
}

func (t *Task) run(msg *message.Message) *errcode.Error {
	for _, v := range t.plugins {
		t.logger.Debug("plugin.Preproc", "plugin", v.Name())
		t.subLogger = t.logger.With("plugin", v.Name(), "fn", "Preproc")
		code, err := v.Preproc(msg, t)
		if code == zplugin.STOP || !reflectx.IsNil(err) {
			return err
		}
	}

	for _, v := range t.plugins {
		t.logger.Debug("plugin.Process", "plugin", v.Name())
		t.subLogger = t.logger.With("plugin", v.Name(), "fn", "Process")
		code, err := v.Process(msg, t)
		if code == zplugin.STOP || !reflectx.IsNil(err) {
			return err
		}
	}

	for _, v := range t.plugins {
		t.logger.Debug("plugin.Postproc", "plugin", v.Name())
		t.subLogger = t.logger.With("plugin", v.Name(), "fn", "Postproc")
		code, err := v.Postproc(msg, t)
		if code == zplugin.STOP || !reflectx.IsNil(err) {
			return err
		}
	}
	return nil
}
