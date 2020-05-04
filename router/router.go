package router

import (
	"api-gin-web/controller/sd"
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
	g.POST("/logout", authMiddleware.LogOutHandler)
	//g.POST("/login", authMiddleware.LoginHandler)

	auth := g.Group("/api/v1")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/dashboard", Dashboard)
		auth.POST("/sysuser", user.LogOut) // 未解决退出登录问题
	}
	// Refresh time can be longer than token timeout
	//g.GET("/refresh_token", authMiddleware.RefreshHandler)

	return g
}

func Dashboard(c *gin.Context) {

	var user = make(map[string]interface{})
	user["login_name"] = "admin"
	user["user_id"] = 1
	user["user_name"] = "管理员"
	user["dept_id"] = 1

	var cmenuList = make(map[string]interface{})
	cmenuList["children"] = nil
	cmenuList["parent_id"] = 1
	cmenuList["title"] = "用户管理"
	cmenuList["name"] = "Sysuser"
	cmenuList["icon"] = "user"
	cmenuList["order_num"] = 1
	cmenuList["id"] = 4
	cmenuList["path"] = "sysuser"
	cmenuList["component"] = "sysuser/index"

	var lista = make([]interface{}, 1)
	lista[0] = cmenuList

	var menuList = make(map[string]interface{})
	menuList["children"] = lista
	menuList["parent_id"] = 1
	menuList["name"] = "Upms"
	menuList["title"] = "权限管理"
	menuList["icon"] = "example"
	menuList["order_num"] = 1
	menuList["id"] = 4
	menuList["path"] = "/upms"
	menuList["component"] = "Layout"

	var list = make([]interface{}, 1)
	list[0] = menuList
	var data = make(map[string]interface{})
	data["user"] = user
	data["menuList"] = list

	var r = make(map[string]interface{})
	r["code"] = 200
	r["data"] = data

	c.JSON(200, r)
}
