package ports

import (
	"WB_L0/internal/service"
	"github.com/gin-gonic/gin"
)

func Router(r gin.IRouter, s service.Service) {
	r.GET("/orders/:uuid", getOrder(s))
}
