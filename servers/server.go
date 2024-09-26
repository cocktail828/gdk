package servers

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"

	"github.com/cocktail828/gdk/v1/errcode"
)

type Server interface {
	Name() string
	Init(data []byte, logger *slog.Logger) *errcode.Error
	Run(ctx context.Context) *errcode.Error
}

type serverIMPL struct {
	initialized atomic.Bool
	server      Server
}

func (srv *serverIMPL) Name() string {
	return srv.server.Name()
}

func (srv *serverIMPL) Init(data []byte, logger *slog.Logger) *errcode.Error {
	if srv.initialized.CompareAndSwap(false, true) {
		return srv.server.Init(data, logger)
	}
	return errcode.CodeRepeatInit.WithMessagef("server:%v hash already been inited", srv.Name())
}

func (srv *serverIMPL) Run(ctx context.Context) *errcode.Error {
	return srv.server.Run(ctx)
}

var (
	providerMu  = sync.Mutex{}
	providerMap = map[string]*serverIMPL{}
)

func Lookup(name string) Server {
	providerMu.Lock()
	defer providerMu.Unlock()
	return providerMap[name]
}

func Register(srv Server) {
	providerMu.Lock()
	defer providerMu.Unlock()
	if srv == nil {
		return
	}
	providerMap[srv.Name()] = &serverIMPL{server: srv}
}
