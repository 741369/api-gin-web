package main

import (
	"context"
	"github.com/741369/go_utils/log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"api-gin-web/router"

	_ "api-gin-web/init"
	"api-gin-web/router/prome"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title api-gin-web API
// @version 1.0.5
// @description 后台接口

// @contact.name lg1024
// @contact.url http://lg1024.com
// @contact.email 1@lg1024.com

// @securityDefinitions.apikey HeaderAuthorization
// @in header
// @name Authorization

// @schemes http https
// @host 127.0.0.1:8060
// @BasePath /
func main() {

	// Create the gin engine
	g := gin.New()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))
	//prome.InitMetrics()
	prome.InitPrometheus()

	//  init routers
	router.Load(g)

	// Start to listening the incoming requests
	log.Infof(nil, "Start to listening the incoming requests on http address: %s", viper.GetString("addr"))

	srv := &http.Server{
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		Addr:           viper.GetString("addr"),
		Handler:        g,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(nil, err, "listen: %s\n", err)
		}
	}()
	//go log.Info(nil, srv.ListenAndServe().Error())

	log.Info(nil, "Enter Control + C Shutdown Server")
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info(nil, "Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(nil, "Server Shutdown:", err)
	}
	log.Info(nil, "Server exiting")
}
