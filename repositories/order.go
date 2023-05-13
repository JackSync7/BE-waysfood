package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(Order models.Order) (models.Order, error)
	DeleteOrder(Order models.Order, ID int) (models.Order, error)
	DeleteAllOrder(IDB int) (models.Order, error)
	FindOrder() ([]models.Order, error)
	GetOrderByUserSeller(IDB int, IDS int) (models.Order, error)
	GetOrderByUserProduct(IDB int, IDS int) (models.Order, error)
	GetOrderbyUser(ID int) ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	UpdateOrder(Order models.Order) (models.Order, error)
	GetByOrderProduct(IDB int, IDS int) (models.Order, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindOrder() ([]models.Order, error) {
	var Order []models.Order
	err := r.db.Preload("Buyer").Preload("Seller").Find(&Order).Error

	return Order, err
}
func (r *repository) GetOrderbyUser(ID int) ([]models.Order, error) {
	var Order []models.Order
	err := r.db.Where("buyer_id = ?", ID).Preload("Buyer").Preload("Product").Preload("Seller").Find(&Order).Error

	return Order, err
}

func (r *repository) GetOrderByUserSeller(IDB int, IDS int) (models.Order, error) {
	var Order models.Order
	err := r.db.Where("buyer_id = ? AND seller_id != ?", IDB, IDS).First(&Order).Error

	return Order, err
}
func (r *repository) GetOrderByUserProduct(IDB int, IDP int) (models.Order, error) {
	var Order models.Order
	err := r.db.Where("buyer_id = ? AND product_id = ?", IDB, IDP).First(&Order).Error

	return Order, err
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var Order models.Order
	err := r.db.First(&Order, ID).Error

	return Order, err
}

func (r *repository) UpdateOrder(Order models.Order) (models.Order, error) {
	err := r.db.Save(&Order).Error
	return Order, err
}

func (r *repository) GetByOrderProduct(IDB int, IDP int) (models.Order, error) {
	var Order models.Order
	err := r.db.First(&Order, IDB, IDP).Error

	return Order, err
}

func (r *repository) CreateOrder(Order models.Order) (models.Order, error) {
	err := r.db.Preload("Buyer").Create(&Order).Error

	return Order, err
}

func (r *repository) DeleteOrder(Order models.Order, ID int) (models.Order, error) {
	err := r.db.Delete(&Order, ID).Scan(&Order).Error

	return Order, err
}
func (r *repository) DeleteAllOrder(ID int) (models.Order, error) {
	var Order models.Order
	err := r.db.Where("buyer_id = ?", ID).Delete(&Order).Error

	return Order, err
}
