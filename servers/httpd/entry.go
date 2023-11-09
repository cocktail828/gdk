package httpd

import (
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/cocktail828/go-tools/netx/httpx"
	"github.com/cocktail828/go-tools/z"
)

func Start(args ...string) {
	handler := Server{}
	z.Must(handler.Init(args...))
	srv := &httpx.Server{
		Addr:    args[0],
		Handler: handler.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("http server stop with err: %v", err)
		}
	}()
	srv.WaitForSignal(time.Second*3, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
}
