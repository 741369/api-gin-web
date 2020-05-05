package router

import (
	"api-gin-web/controller/sd"
	"api-gin-web/controller/system"
	"api-gin-web/controller/user"
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

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	if err != nil {
		panic("JWT init error" + err.Error())
	}

	g.POST("/login", authMiddleware.LoginHandler)
	g.POST("/logout", user.Logout) // 未解决退出登录问题

	apiv1 := g.Group("/api/v1")
	{
		apiv1.GET("/dashboard", system.Dashboard)       // 获取首页dashboard
		apiv1.GET("/monitor/server", system.ServerInfo) // 获取服务器信息
		apiv1.GET("/captcha", system.GenerateCaptcha)   // 生成图片验证码
		apiv1.GET("/db/tables/page", system.GetDBTableList)
		apiv1.GET("/db/columns/page", system.GetDBColumnList)

		apiv1.GET("/sys/tables/page", system.GetSysTableList)
		apiv1.POST("/sys/tables/info", system.InsertSysTable)
		apiv1.PUT("/sys/tables/info", system.UpdateSysTable)
		apiv1.DELETE("/sys/tables/info/:tableId", system.DeleteSysTables)
		apiv1.GET("/sys/tables/info/:tableId", system.GetSysTables)
		apiv1.GET("/gen/preview/:tableId", system.Preview)
		apiv1.GET("/menuTreeselect", system.GetMenuTreeelect)
		apiv1.GET("/dict/databytype/:dictType", system.GetDictDataByDictType)
	}

	auth := g.Group("/api/v1")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/sysuser", user.InsertSysUser)
	}

	return g
}
