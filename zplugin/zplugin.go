package zplugin

import (
	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/message"
	"github.com/cocktail828/go-tools/messagepb"
	"golang.org/x/exp/slog"
)

type State int

const (
	STOP State = iota
	CONT
)

const New = "New" // plugin entry
type Entry = func() ZPlugin

type ZPlugin interface {
	Name() string
	Interest(msg *message.Message) bool
	Init(cfgs map[string][]byte) *errcode.Error
	Preproc(msg *message.Message, tools Tools) (State, *errcode.Error)
	Process(msg *message.Message, tools Tools) (State, *errcode.Error)
	Postproc(msg *message.Message, tools Tools) (State, *errcode.Error)
}

type MessageParser interface {
	Parser(msg *messagepb.Message) (*message.Parsed, error)
	Sub() string
}

type Tools interface {
	Logger() *slog.Logger
	SendBack([]byte)
}
