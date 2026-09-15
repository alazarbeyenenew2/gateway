package initiator

import (
	"github.com/alazarbeyenenew2/gateway/internal/glue/auth"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func initRoute(
	group *gin.RouterGroup,
	handler *Handler,
	log *zap.Logger,

) {
	auth.Init(group, log, handler.AuthHandler)
}
