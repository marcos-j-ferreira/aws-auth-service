package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
	"log"
)

// Variavel globla para a chave do JWT
var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims representa as informações que srão armazenadas no JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: 	jwt.NewNumericDate(expirationTime),
			IssuedAt:	jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}


func TokenDefin(){
	if len(JwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY environment variable not set")
	}
}

