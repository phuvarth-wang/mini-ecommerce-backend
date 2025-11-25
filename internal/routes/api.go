package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuvarth-wang/mini-ecommerce-backend/internal/handlers"
	"github.com/phuvarth-wang/mini-ecommerce-backend/internal/middleware"
	"github.com/phuvarth-wang/mini-ecommerce-backend/internal/repositories"
	"github.com/phuvarth-wang/mini-ecommerce-backend/internal/services"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// ------------------------------
	// Product setup
	// ------------------------------
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Public routes (everyone)
	r.GET("/products", productHandler.GetProducts)
	r.GET("/products/:id", productHandler.GetProduct)

	// Admin-only routes
	admin := r.Group("/admin")
	admin.Use(middleware.AdminOnly())
	{
		admin.POST("/products", productHandler.CreateProduct)
		admin.PUT("/products/:id", productHandler.UpdateProduct)
		admin.DELETE("/products/:id", productHandler.DeleteProduct)
	}
}
