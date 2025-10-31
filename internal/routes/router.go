package routes

import (
	"ci/cd/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"

	"github.com/joho/godotenv"
	"os"
	"ci/cd/internal/auth"
	"fmt"

	"strings"
)

func SetupRouter() *gin.Engine {

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router.Use(gin.Recovery())

	_ = godotenv.Load(".env")

	ip := os.Getenv("IP_FRONT")
	var (
		ip2 			string
		ip3 			string
		ip4Register 	string
		ip5Login 		string
		ip6Users 		string
	)

	if !strings.HasPrefix(ip, "http://") && !strings.HasPrefix(ip, "https://") {
		
		ip2 		= fmt.Sprintf("http://%s:8080", ip)
		ip3 		= fmt.Sprintf("http://%s", ip)
		ip4Register	= fmt.Sprintf("http://%s/login.html", ip)
		ip5Login	= fmt.Sprintf("http://%s/register.html", ip)
		ip6Users	= fmt.Sprintf("http://%s/users.html", ip)
		ip 			= fmt.Sprintf("http://%s:80",ip)

	}
	//fmt.Println(ip, ip2)
	// Configuração de CORS - somente para meu ip
	router.Use(cors.New(cors.Config{
		AllowOrigins:		[]string{

			ip,
			ip2,
			ip3,
			ip4Register,
			ip5Login,
			ip6Users,
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

	
	auth.Init_Connection_db()


	// Rotas para autenticação
	auth_v1 := router.Group("v1")
	{
		auth_v1.POST("/register", auth.Register)
		auth_v1.POST("/login", auth.Login)

		//auth_v1.GET("/users", auth.GetAllUsers)
	}
	//auth.TokenDefin()
	protected := router.Group("protected")
	protected.Use(auth.AuthMiddleware())	
	{
		protected.GET("/profile", handler.Profile)
		protected.GET("/users", auth.GetAllUsers)
	}

	users := router.Group("users")
	{
		users.GET("/", auth.GetAllUsers)
	}


	
	return router
}
