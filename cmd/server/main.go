package main

import (
	//"github.com/gin-gonic/gin"
	"ci/cd/internal/routes"
	//"ci/cd/internal/handler"
	"fmt"
)


func main() {
	//gin.SetMode(gin.ReleaseMode)

	router := routes.SetupRouter()

	//r := gin.New()
	//router.Use(gin.Recovery())
	//router.GET("/", handler.HandlerHello)
	//router.GET("/ola", handler.HandlerOla)
	//router.GET("/deploy", handler.HandlerTesteDeploy)

	fmt.Println("Servidor ouvindo na porta 8080...")

	router.Run(":8080")
}