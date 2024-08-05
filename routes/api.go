package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRouter(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("/", func(r *gin.Context) {
			r.JSON(http.StatusOK, gin.H{
				"hello": "World",
			})
		})
	}
}
