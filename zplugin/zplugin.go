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
	Init(conf string) *errcode.Error
	Name() string
	Preproc(msg *message.Message, tools Tools) (State, *errcode.Error)
	Process(msg *message.Message, tools Tools) (State, *errcode.Error)
	Postproc(msg *message.Message, tools Tools) (State, *errcode.Error)
	Interest(msg *message.Message) bool
}

type Tools interface {
	Logger() *slog.Logger
	SendBack([]byte)
}
