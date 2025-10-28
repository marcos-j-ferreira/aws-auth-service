package handler


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func HandlerHello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"mensagem":"Hello, world!!"})
}

func HandlerOla(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mensagem":"Olá mundo!!"})
}

func HandlerTesteDeploy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mensagem":"Deploy está funcionando com sucesso"})
}

// profile é o handler para a rota protegida /protected/profile
func Profile(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message":"Welcome to your profile", "username":username})
}