package auth

import (
	//"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	Username		 string	`json:"username" binding: "required"`
	Password		 string	`json:"password" binding: "required"` 
}

//In-memory storoge para user
var users = make(map[string]User)
var mutex = &sync.RWMutex{}

func Register(c *gin.Context){
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// 1 Verificar se o usuario já existe

	if _, exists := users[input.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error":"Username already taken"})
		return
	}

	// 2 Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to hash password"})
		return
	}

	newUser := User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	users[input.Username] = newUser
	c.JSON(http.StatusCreated, gin.H{"message":"User registered successfully"})

	// Test rápido
	//curl -X POST http://localhost:8080/v1/register -H "Content-Type: application/json" -d '{"username":"testuser", "password":"testpassword"}'

}

func Login(c *gin.Context) {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	mutex.RLock()
	user, exists := users[input.Username]
	mutex.RUnlock()

	// Verifica se usuario existe
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid credentials"})
		return
	}

	// Compara a senha fornecida com a senha hashada
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid credentials"})
		return
	}

	tokenString, err := GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failed to geneate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})

	//curl -X POST http://localhost:8080/v1/register -H "Content-Type: application/json" -d '{"username":"testuser", "password":"testpassword"}'
}

