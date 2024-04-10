package zplugin

import (
	"log"
	"sync"

	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/zplugin/consts"
)

var (
	repository sync.Map
)

func Register(p ZPlugin) {
	if p != nil {
		log.Println("register plugin", p.Name())
		repository.Store(p.Name(), p)
	}
}

func Traverse(f func(p ZPlugin) *errcode.Error, handlers ...string) {
	for _, handler := range handlers {
		if val, ok := repository.Load(handler); ok {
			if err := f(val.(ZPlugin)); err != nil {
				return
			}
		}
	}
}

type ZPlugin interface {
	Name() string
	Init(cfg []byte) *errcode.Error
	Interest(ctx *Context) bool
	Prepare(ctx *Context) (consts.State, *errcode.Error)
	Process(ctx *Context) (consts.State, *errcode.Error)
	Cleanup(ctx *Context) (consts.State, *errcode.Error)
}
