package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phuvarth-wang/mini-ecommerce-backend/internal/database"
	"github.com/phuvarth-wang/mini-ecommerce-backend/internal/routes"
)

func main() {
	godotenv.Load()
	db := database.Connect()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		log.Fatal("Missing SUPABASE_URL in .env")
	}

	log.Println("Using Supabase HS256 Auth Mode")
	routes.RegisterRoutes(r, db)

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
