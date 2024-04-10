package httpd

import (
	"net/http"

	"github.com/cocktail828/go-kits/pkg/logger"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type Server struct {
	*gin.Engine
	state  state
	logger logger.Logger
}

func (srv *Server) Init(logger logger.Logger, args ...string) error {
	gin.SetMode(gin.ReleaseMode)
	srv.logger = logger
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
