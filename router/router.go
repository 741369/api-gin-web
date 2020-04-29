package router

import (
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
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 监控服务器性能API
	//g.GET("/sd/health", monitor.HealthCheck)

	return g
}
