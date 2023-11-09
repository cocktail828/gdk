package httpd

import (
	"net/http"

	server "github.com/cocktail828/gdk/v1/servers"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type Server struct {
	*gin.Engine
	state  state
	logger *slog.Logger
}

func (srv *Server) Init(args ...string) error {
	gin.SetMode(gin.ReleaseMode)
	srv.logger = server.NewLogger("gdk-http.log")
	srv.Engine = gin.New()

	srv.Use(gin.Recovery())
	srv.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, srv.state.String())
	})
	return nil
}

func (srv *Server) Handler() http.Handler {
	return srv.Engine
}

func (srv *Server) RegisterHandlers(addr string) {
	// v1 := g.Group("/v1")
}
