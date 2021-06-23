package main

import (
	"context"
	"courses-gin/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := gin.Default()

	// 加载网站静态资源
	r.StaticFS("/static", http.Dir("D:/go_progrem/courses-gin/website/static"))
	r.LoadHTMLGlob("website/tpl/*")

	// 加载路由
	router.LoadRouter(r)

	// 服务器配置
	srv := &http.Server{
		Addr:         ":80",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅Shutdown（或重启）服务
	// 5秒后优雅Shutdown服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) //syscall.SIGKILL
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
