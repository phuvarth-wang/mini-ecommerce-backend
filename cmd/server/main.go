package main

import (
	"log"

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
	routes.RegisterRoutes(r, db)

	log.Println("Sever running on http://localhost:8080")
	r.Run(":8080")
}