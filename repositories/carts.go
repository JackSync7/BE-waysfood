package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type CartsRepository interface {
	CreateCarts(Carts models.Carts) (models.Carts, error)
	DeleteCarts(Carts models.Carts, ID int) (models.Carts, error)
	FindCarts() ([]models.Carts, error)
	GetCarts(ID int) (models.Carts, error)
}

func RepositoryCarts(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindCarts() ([]models.Carts, error) {
	var carts []models.Carts
	err := r.db.Find(&carts).Error

	return carts, err
}

func (r *repository) GetCarts(ID int) (models.Carts, error) {
	var carts models.Carts
	err := r.db.First(&carts, ID).Error

	return carts, err
}

func (r *repository) CreateCarts(Carts models.Carts) (models.Carts, error) {
	err := r.db.Create(&Carts).Error

	return Carts, err
}

func (r *repository) DeleteCarts(Carts models.Carts, ID int) (models.Carts, error) {
	err := r.db.Delete(&Carts, ID).Scan(&Carts).Error

	return Carts, err
}
