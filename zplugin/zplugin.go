package zplugin

import (
	"github.com/cocktail828/gdk/v1/message"
	"github.com/cocktail828/go-tools/z/errcode"
	"golang.org/x/exp/slog"
)

type State int

const (
	STOP State = iota
	CONT
)

const New = "New" // plugin entry

type ZPlugin interface {
	Name() string
	Interest(msg *message.Message) bool
	Init(cfgs map[string][]byte) *errcode.Error
	Preproc(msg *message.Message, tools Tools) (State, *errcode.Error)
	Process(msg *message.Message, tools Tools) (State, *errcode.Error)
	Postproc(msg *message.Message, tools Tools) (State, *errcode.Error)
}

type Tools interface {
	Logger() *slog.Logger
	SendBack([]byte)
}
