package main

import (
	"Quortle/internal/api"
	models "Quortle/internal/models"
	"Quortle/internal/repository"
	"Quortle/internal/server"
	"Quortle/internal/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load .env
	LoadEnv()

	if GetEnv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	ConnectDatabase()
	DB.Exec("CREATE SCHEMA IF NOT EXISTS quortle")
	if err := DB.AutoMigrate(models.AllModels...); err != nil {
		log.Fatal("Migration failed:", err)
	}
	syncWordsToDB("words.txt")

	repo := &repository.WordRepository{FilePath: "words.txt"}

	wordSvc := services.NewWordService(repo)
	userSvc := services.NewUserService(DB)

	handler := api.NewHandler(wordSvc, userSvc)
	router := handler.Routes()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	s := server.NewServer(router, GetEnv("DOMAIN"))
	s.Start()
}
