package model

type Payment struct {
	ID          int  `json:"id"`
	Category_Id int  `json:"category_id"` // Required field
	Date        *int `json:"date"`        // Required field

}

type PostPayment struct {
	Category_Id int  `json:"category_id"` // Required field
	Date        *int `json:"date"`
}

type PaymentUri struct {
	ID int `uri:"id" binding:"required,number"`
}
