package gin

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinServer(t *testing.T) {
	r := gin.Default()
	r.GET("/test/:err", func(c *gin.Context) {
		err := c.Param("err")
		if err == "true" {
			c.JSON(http.StatusBadRequest, "error")
			return
		}
		c.JSON(http.StatusOK, "ok")
	})
	r.Run(":8080")
}
