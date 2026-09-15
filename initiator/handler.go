package initiator

import (
	"github.com/alazarbeyenenew2/gateway/internal/constants/dto"
	"github.com/alazarbeyenenew2/gateway/internal/handler"
	"github.com/alazarbeyenenew2/gateway/internal/handler/auth"
	"go.uber.org/zap"
)

type Handler struct {
	AuthHandler handler.Auth
}

func InitHandler(logger *zap.Logger, client dto.GRPCClients) *Handler {
	return &Handler{
		AuthHandler: auth.Init(logger, client.Auth),
	}
}
