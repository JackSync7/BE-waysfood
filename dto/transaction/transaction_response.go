package transactiondto

type TransactionDeleteResponse struct {
	ID int `json:"id" gorm:"primary_key:auto_increment"`
}
