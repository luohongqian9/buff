package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
func initServer(addr string, router *gin.Engine) server {
	s := endless.NewServer(addr, router)
	s.ReadHeaderTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	return s
}
*/

func initServer(addr string, router *gin.Engine) server {
	return &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
