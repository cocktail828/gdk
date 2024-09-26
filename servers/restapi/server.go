package restapi

import (
	"context"
	"log/slog"

	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/servers"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	servers.Register(&Service{})
}

type Service struct {
	engine *gin.Engine
	logger *slog.Logger
}

func (svc *Service) Name() string {
	return "restapi"
}

func (svc *Service) Init(data []byte, logger *slog.Logger) *errcode.Error {
	svc.logger = logger
	svc.engine = gin.New()

	svc.engine.Use(gin.Recovery())
	return nil
}

func (svc *Service) Run(ctx context.Context) *errcode.Error {
	return nil
}
