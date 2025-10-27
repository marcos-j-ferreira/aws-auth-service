package routes

import (
	"ci/cd/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Configuração de CORS - somente para meu ip
	router.Use(cors.New(cors.Config{
		AllowOrigins:		[]string{
			"http://13.218.89.228",
			"http://13.218.89.228:8080",

			//"http://localhost",
			//"http://localhost:8080",
			//"http://192.168.1.111:80/index.html",

			//"http://127.0.0.1:3000",    /// unica que funcionou
			
			//"http://192.168.1.111:5500", 
			//"http://192.168.1.111:3000",  
			 },
		AllowMethods:		[]string{"GET","POST","PUT","DELETE", "OPTIONS"},
		AllowHeaders:		[]string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:		[]string{"Content-Length"},
		AllowCredentials:	true,
		MaxAge:				12* time.Hour,
	}))

	// Rotas
	router.GET("/", handler.HandlerHello)
	router.GET("/ola", handler.HandlerOla)
	router.GET("/deploy", handler.HandlerTesteDeploy)

	return router
}
