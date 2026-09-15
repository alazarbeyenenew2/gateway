package auth

import (
	"net/http"

	authGenerated "github.com/alazarbeyenenew2/common/generated/auth/auth"
	"github.com/alazarbeyenenew2/gateway/internal/constants/dto"
	"github.com/alazarbeyenenew2/gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type auth struct {
	logger *zap.Logger
	client authGenerated.AuthServiceClient
}

func Init(logger *zap.Logger, client authGenerated.AuthServiceClient) handler.Auth {

	return &auth{
		logger: logger,
		client: client,
	}
}

func (a *auth) Login(c *gin.Context) {
	var loginReq dto.AuthRequest
	if err := c.ShouldBind(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "invalid request",
		})
	}

	resp, err := a.client.AuthenticateUser(c, &authGenerated.AuthenticateUserRequest{
		Username:      loginReq.Username,
		Password:      loginReq.Password,
		XforwardedFor: loginReq.XforwardedFor,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": dto.AuthResponse{
			Token: resp.Token,
		},
	})
}
