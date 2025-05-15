package model

type Transactionreport struct {
	ID          int  `json:"id"`
	Customer_Id int  `json:"customer_id"` // Required field
	Order_Id    *int `json:"order_id"`    // Required field
	Product_Id  *int `json:"product_id"`
	Payment_Id  *int `json:"payment_id"`
}

type PostTransactionreport struct {
	Customer_Id int  `json:"customer_id"` // Required field
	Order_Id    *int `json:"order_id"`    // Required field
	Product_Id  *int `json:"product_id"`
	Payment_Id  *int `json:"payment_id"`
}

type TransactionreportUri struct {
	ID int `uri:"id" binding:"required,number"`
}
