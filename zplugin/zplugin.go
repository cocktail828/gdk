package zplugin

import "github.com/cocktail828/gdk/v1/message/messagepb"

const (
	STOP = iota
	CONT = iota
)

const New = "New" // plugin entry

type ZPlugin interface {
	Init(conf string) error
	Name() string
	Prepare(message *messagepb.Message, tools Tools) (int, error)
	Do(message *messagepb.Message, tools Tools) (int, error)
	WindingUp(message *messagepb.Message, tools Tools) (int, error)
	Interest(mess *messagepb.Message) bool
}

// type Tools interface {
// 	Plugins() []ZPlugin
// 	Span() *utils.Span
// 	Logger() *utils.Logger
// 	Tools() *xsf.ToolBox
// 	Skip(string string)
// 	SendBack([]byte)
// }
