package model

type Order struct {
	ID          int  `json:"id"`
	Customer_Id int  `json:"customer_id"` // Required field
	Date        *int `json:"date"`        // Required field

}

type PostOrder struct {
	Customer_Id int  `json:"customer_id"` // Required field
	Date        *int `json:"date"`
}

type OrderUri struct {
	ID int `uri:"id" binding:"required,number"`
}
