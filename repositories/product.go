package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
	FindProductPartner(ID int) ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(Product models.Product) (models.Product, error)
	UpdateProduct(Product models.Product) (models.Product, error)
	DeleteProduct(Product models.Product, ID int) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProduct() ([]models.Product, error) {
	var Products []models.Product
	err := r.db.Preload("User").Find(&Products).Error

	return Products, err
}

func (r *repository) FindProductPartner(ID int) ([]models.Product, error) {
	var Products []models.Product
	err := r.db.Preload("User").Find(&Products, "user_id = ?", ID).Error

	return Products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var Products models.Product
	err := r.db.Preload("Category").First(&Products, ID).Error

	return Products, err
}

func (r *repository) CreateProduct(Product models.Product) (models.Product, error) {
	err := r.db.Create(&Product).Error

	return Product, err
}

func (r *repository) UpdateProduct(Product models.Product) (models.Product, error) {
	err := r.db.Save(&Product).Error
	return Product, err
}

func (r *repository) DeleteProduct(Product models.Product, ID int) (models.Product, error) {
	err := r.db.Delete(&Product, ID).Scan(&Product).Error

	return Product, err
}
