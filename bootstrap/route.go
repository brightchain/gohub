package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	//中间件
	registerGlobalMiddleWare(router)

	//api路由
	routes.RegisterAPIRouter(router)

	//404页面路由
	Set404Handler(router)
}

func registerGlobalMiddleWare(route *gin.Engine) {
	route.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func Set404Handler(route *gin.Engine) {
	route.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
