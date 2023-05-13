package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	GetUserTransaction(ID int) ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(Transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
	DeleteTransaction(Transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var Transaction []models.Transaction
	err := r.db.Preload("Cart.Product").Preload("Buyer").Preload("Seller").Find(&Transaction).Error

	return Transaction, err
}
func (r *repository) GetUserTransaction(ID int) ([]models.Transaction, error) {
	var Transaction []models.Transaction
	err := r.db.Where("buyer_id = ?", ID).Preload("Cart.Product").Preload("Buyer").Preload("Seller").Find(&Transaction).Error

	return Transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var Transaction models.Transaction
	err := r.db.Preload("User").First(&Transaction, ID).Error

	return Transaction, err
}

func (r *repository) CreateTransaction(Transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Film").Create(&Transaction).Error

	return Transaction, err
}

func (r *repository) DeleteTransaction(Transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&Transaction, ID).Scan(&Transaction).Error

	return Transaction, err
}

func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, orderId)

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}
