package web

type MidtransRequest struct {
	UserID   int    `json:"user_id" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
	ItemID   string `json:"item_id" binding:"required"`
	ItemName string `json:"item_name" binding:"required"`
}
