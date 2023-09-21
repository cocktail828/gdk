package proxy

import (
	"bytes"

	"github.com/cocktail828/gdk/v1/logger"
	"github.com/cocktail828/gdk/v1/message/messagepb"
	"github.com/cocktail828/gdk/v1/responsibe_chain"
	"github.com/cocktail828/gdk/v1/zplugin"
)

type taskOpt func(*Task)

func WithConf(conf interface{}) taskOpt {
	return func(t *Task) {
		t.params["conf"] = conf
	}
}

func WithChain(chain []zplugin.ZPlugin) taskOpt {
	return func(t *Task) {
		t.chain = chain
	}
}

func WithSkip(skipPlugins map[string]bool) taskOpt {
	return func(t *Task) {
		chain := []zplugin.ZPlugin{}
		for _, p := range t.chain {
			if skipPlugins[p.Name()] {
				chain = append(chain, p)
			}
		}
		t.chain = chain
	}
}

func WithBuf(buf bytes.Buffer) taskOpt {
	return func(t *Task) {
		t.buf = buf
	}
}

type Task struct {
	logger logger.Logger
	params map[string]interface{}
	chain  []zplugin.ZPlugin
	buf    bytes.Buffer
}

func NewTask(opts ...taskOpt) (*Task, error) {
	taskInst := &Task{logger: logger.Default()}
	for _, opt := range opts {
		opt(taskInst)
	}

	return taskInst, nil
}

func (t *Task) run(mess *messagepb.Message) (int, error) {
	for _, v := range t.chain {
		logger.Debugln("task.run", "op", "Prepare", "filter", v.Name())
		code, err := v.Preproc(mess, t)
		if err != nil {
			return code, err
		}

		if code == zplugin.STOP {
			return code, err
		}
	}

	for _, v := range t.chain {
		logger.Debugln("task.run", "op", "Do", "filter", v.Name())
		code, err := v.Process(mess, t)
		if err != nil {
			return code, err
		}

		if code == zplugin.STOP {
			return code, err
		}
	}

	for _, v := range t.chain {
		logger.Debugln("task.run", "op", "WindingUp", "filter", v.Name())
		code, err := v.Postproc(mess, t)
		if err != nil {
			return code, err
		}

		if code == zplugin.STOP {
			return code, err
		}
	}

	return 0, nil
}

func (t *Task) Run(mess *messagepb.Message) error {
	zplugin.Plugins(func(res responsibe_chain.Handler) {
		v := res.(zplugin.ZPlugin)
		if v.Interest(mess) {
			t.chain = append(t.chain, v)
		}
	})

	_, err := t.run(mess)
	return err
}

func (t *Task) Logger() logger.Logger {
	return t.logger
}

func (t *Task) SendBack(d []byte) {
	t.buf.Write(d)
}
