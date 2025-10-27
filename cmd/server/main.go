package main

import (
	"github.com/gin-gonic/gin"
	"ci/cd/internal/handler"
	"fmt"
)


func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/", handler.HandlerHello)
	r.GET("/ola", handler.HandlerOla)
	r.GET("/deploy", handler.HandlerTesteDeploy)

	fmt.Println("Servidor ouvindo na porta 8080...")

	r.Run(":8080")
}