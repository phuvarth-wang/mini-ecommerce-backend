package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phuvarth-wang/mini-ecommerce-backend/internal/database"
	"github.com/phuvarth-wang/mini-ecommerce-backend/internal/routes"
	"github.com/gin-contrib/cors"
)

func main() {
	godotenv.Load()
	db := database.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.SetTrustedProxies(nil)

	log.Println("Using Supabase HS256 Auth Mode")
	routes.RegisterRoutes(r, db)

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
