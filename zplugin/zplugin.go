package zplugin

import (
	"github.com/cocktail828/gdk/v1/logger"
	"github.com/cocktail828/gdk/v1/message/messagepb"
)

const (
	STOP = iota
	CONT = iota
)

const New = "New" // plugin entry

type ZPlugin interface {
	Init(conf string) error
	Name() string
	Preproc(message *messagepb.Message, tools Tools) (int, error)
	Process(message *messagepb.Message, tools Tools) (int, error)
	Postproc(message *messagepb.Message, tools Tools) (int, error)
	Interest(mess *messagepb.Message) bool
}

type Tools interface {
	Logger() logger.Logger
	SendBack([]byte)
}
