package repositories

import (
    "github.com/phuvarth-wang/mini-ecommerce-backend/internal/models"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{
        db: db,
    }
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
    var products []models.Product
    err := r.db.Preload("Category").Find(&products).Error
    return products, err
}

func (r *ProductRepository) FindByID(id uuid.UUID) (*models.Product, error) {
    var product models.Product
    err := r.db.Preload("Category").First(&product, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepository) Create(product *models.Product) error {
    return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
    return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&models.Product{}, "id = ?", id).Error
}
