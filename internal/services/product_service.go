package services

import (
    "github.com/google/uuid"
    "github.com/phuvarth-wang/mini-ecommerce-backend/internal/models"
    "github.com/phuvarth-wang/mini-ecommerce-backend/internal/repositories"
)

type ProductService struct {
    repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
    return &ProductService{
        repo: repo,
    }
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
    return s.repo.FindAll()
}

func (s *ProductService) GetProduct(id uuid.UUID) (*models.Product, error) {
    return s.repo.FindByID(id)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
    return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
    return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uuid.UUID) error {
    return s.repo.Delete(id)
}