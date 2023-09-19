package proxy

import (
	"bytes"

	"github.com/cocktail828/gdk/v1/logger"
	"github.com/cocktail828/gdk/v1/message/messagepb"
	"github.com/cocktail828/gdk/v1/zplugin"
)

type taskOpt func(*Task)

func WithConf(conf interface{}) taskOpt {
	return func(in *Task) {
		in.params["conf"] = conf
	}
}

func WithChain(chain []zplugin.ZPlugin) taskOpt {
	return func(in *Task) {
		in.chain = chain
	}
}

func WithSkip(skipPlugins map[string]bool) taskOpt {
	return func(in *Task) {
		in.skipPlugins = skipPlugins
	}
}

func WithBuf(buf bytes.Buffer) taskOpt {
	return func(in *Task) {
		in.buf = buf
	}
}

type Task struct {
	params      map[string]interface{}
	chain       []zplugin.ZPlugin
	skipPlugins map[string]bool
	buf         bytes.Buffer
}

func NewTask(opts ...taskOpt) (*Task, error) {
	taskInst := &Task{
		skipPlugins: make(map[string]bool),
	}
	for _, opt := range opts {
		opt(taskInst)
	}

	return taskInst, nil
}

func (t *Task) run(mess *messagepb.Message, span *utils.Span) (int, error) {
	for _, v := range t.chain {
		if t.skipPlugins[v.Name()] {
			continue
		}

		logger.Default().Debugw("task.run", "op", "Prepare", "filter", v.Name())
		code, err := v.Prepare(mess, t)
		if err != nil {
			t.tools.Log.Errorw("task.run", "op", "Prepare", "err", err, "sid", t.Sid)
			return code, err
		}

		if code == zplugin.STOP {
			return code, err
		}
	}

	for _, v := range t.chain {
		if t.skipPlugins[v.Name()] {
			continue
		}
		logger.Default().Debugw("task.run", "op", "Do", "filter", v.Name())
		code, err := v.Do(mess, t)
		if err != nil {
			t.tools.Log.Errorw("task.run", "op", "Do", "err", err, "sid", t.Sid)
			return code, err
		}

		if code == zplugin.STOP {
			return code, err
		}
	}

	for _, v := range t.chain {
		if t.skipPlugins[v.Name()] {
			continue
		}

		logger.Default().Debugw("task.run", "op", "WindingUp", "filter", v.Name())
		code, err := v.WindingUp(mess, t)
		if err != nil {
			t.tools.Log.Errorw("task.run", "op", "WindingUp", "err", err, "sid", t.Sid)
			return code, err
		}

		if code == zplugin.STOP {
			return code, err
		}
	}

	for k, v := range mess.Params.Extra {
		span.WithTag(k, v)
	}

	return 0, nil
}

func (t *Task) Run(mess *messagepb.Message) error {
	t.Sid = mess.Params.Sid
	str := ""
	chain := make([]zplugin.ZPlugin, 0)

	logger.Default().Debugw("about to GetFiltersChains")
	chainPreset, err := GetChain("")
	if err == nil {
		return err
	}

	for _, v := range chainPreset {
		if v.Interest(mess) {
			chain = append(chain, v)
			str += v.Name() + ", "
			logger.Default().Debugw("task.Run", "str", str)
		}
	}

	t.span.WithTag("filters: ", str)
	t.chain = chain
	code, err = t.run(mess, t.span)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) Chain() []zplugin.ZPlugin{
	return t.chain
}

func (t *Task) Skip(name string) {
	t.skipPlugins[name] = true
}

func (t *Task) SendBack(d []byte) {
	t.buf.Write(d)
}

func (t *Task) GetBackData() []byte {
	return t.buf.Bytes()
}
