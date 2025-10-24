package handler


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerHello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"mensagem":"Hello, world!!"})
}