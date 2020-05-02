package router

import (
	"api-gin-web/controller/sd"
	"net/http"

	_ "api-gin-web/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api-gin-web/router/middleware"

	recovery "api-gin-web/router/middleware/recover"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// Middleware
	g.Use(gin.RecoveryWithWriter(recovery.NewWriter()))
	g.Use(mw...)
	g.Use(middleware.Logging()) // 中间件，监控所有请求并打印日志
	g.Use(middleware.TraceId()) // trace_id
	g.Use(middleware.CorsMiddleware())

	// 404 Handler
	g.NoRoute(func(context *gin.Context) {
		if context.Request.Method != "OPTIONS" {
			context.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    -999,
				"message": "The incorrect API route.",
				"data":    nil,
			})
		}
	})

	// swagger api docs
	swaggerRouter := g.Group("/swagger")
	swaggerRouter.Use(middleware.PushWhite())
	{
		swaggerRouter.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g.Static("/static", "./static")

	// 监控信息
	svcd := g.Group("/sd")
	{
		svcd.GET("/info", sd.Ping)
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/os", sd.OSCheck)
		svcd.GET("/metrics", gin.WrapH(promhttp.Handler())) // prometheus监控
	}

	//g.POST("/login", authMiddleware.LoginHandler)

	// Refresh time can be longer than token timeout
	//g.GET("/refresh_token", authMiddleware.RefreshHandler)

	return g
}
