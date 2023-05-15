package ports

import (
	"WB_L0/internal/orders"
	"WB_L0/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func getOrder(a service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidStr := c.Param("uuid")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			c.Status(http.StatusBadRequest)
			c.Writer.WriteString(err.Error())
			c.Writer.Flush()
			return
		}
		orderRaw, err := a.GetOrder(uid)

		if err != nil {
			c.Status(http.StatusBadRequest)
			c.Writer.WriteString(err.Error())
			c.Writer.Flush()
			return
		}
		var order orders.OrderModel
		json.Unmarshal(orderRaw.Data, &order)
		c.JSON(http.StatusOK, order)
	}
}
