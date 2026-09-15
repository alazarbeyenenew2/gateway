package handler

import "github.com/gin-gonic/gin"

type Auth interface {
	Login(c *gin.Context)
}
