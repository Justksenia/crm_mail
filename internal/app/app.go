package app

import (
	"log"
	"mail/internal/db"
	"mail/internal/models"
	"mail/internal/server"
	"mail/internal/services"
	"mail/internal/transport/rest"
	"os"

	_"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Run() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error on load env: %s", err)
	}
	db, err := db.NewPostgresDb(
		db.Config{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName: os.Getenv("DB_NAME"),
			SSLMode: "disable",
	})

	if err != nil {
		log.Fatalf("fall DB: %s", err.Error())
	}
	
	// rdb := redis.NewClient(&redis.Options{
	// 		Addr: "10.63.2.143:6379",
	// 		Password: "", 
	// 		DB: 0, 
	// 	})
	
	models := models.NewModel(db)
	services := services.NewService(*models)
	handlers := rest.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error connect")
	}
}
