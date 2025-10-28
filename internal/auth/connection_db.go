package auth

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"strconv"
	"time"
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

// Variavel global para a conexão com o banco de dados
var DB *gorm.DB

func Init_Connection_db() {

	// env
	_ = godotenv.Load(".env")

	host 			:= os.Getenv("DB_HOST")
	port 			:= os.Getenv("DB_PORT")
	user			:= os.Getenv("DB_USER")
	password		:= os.Getenv("DB_PASSWORD")
	dbname			:= os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbname, port)

	var err error

	// Tentar conectar mais de algumas vezes, por validações
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			break
		}

		log.Printf("Failed to connect to databases. Retrying in 5 seconds... Error: %v", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to database after multiple retries: ", err)
	}

	// Migrar o esquema
	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to auto migrate databases: ", err)
	}

	log.Println("\nDatabase connection and migration successfuly\n")
} 
