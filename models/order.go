package models

type Order struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Qty       int                  `json:"qty" `
	BuyerID   int                  `json:"buyer_id" `
	Buyer     UsersProfileResponse `json:"buyer" `
	SellerID  int                  `json:"seller_id" `
	Seller    UsersProfileResponse `json:"seller" `
	ProductID int                  `json:"product_id" gorm:"type: int"`
	Product   Product              `json:"product" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
