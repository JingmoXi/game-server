package http_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
