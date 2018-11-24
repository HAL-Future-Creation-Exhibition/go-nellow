package router

import (
	"net/http"

	"github.com/HAL-Future-Creation-Exhibition/go-nellow/controller"
	"github.com/gin-gonic/gin"
)

func userApiRouter(user *gin.RouterGroup) {
	user.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ようこそ User")
	})
	user.GET("/", controller.UserController.Get)
	user.POST("/", controller.UserController.Create)
}
