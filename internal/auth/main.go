package auth

import (
	//"log"
	"net/http"
	"sync"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	gorm.Model
	Username		 string	`gorm: "uniqueIndex;not null" json:"username" binding:"required"`
	Password		 string	`gorm: "not null" json:"password" binding:"required"`
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

	var existingUser User
	if err := DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error":"Username already taken"})
		return
	}else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Database error"})
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

	if err := DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":" Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"User registered successfully"})

}

func Login(c *gin.Context) {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	var user User

	if err := DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Database error"})
		}
		return
	}

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

}
