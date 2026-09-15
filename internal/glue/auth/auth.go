package auth

import (
	"net/http"

	"github.com/alazarbeyenenew2/gateway/internal/glue/routing"
	"github.com/alazarbeyenenew2/gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(
	group *gin.RouterGroup,
	log *zap.Logger,
	authHandler handler.Auth,
) {

	weatherRoutes := []routing.Route{
		{
			Method:     http.MethodPost,
			Path:       "/login",
			Handler:    authHandler.Login,
			Middleware: []gin.HandlerFunc{},
		},
	}
	routing.RegisterRoute(group, weatherRoutes, log)
}
