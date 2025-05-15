package model

type Delivery struct {
	ID          int  `json:"id"`
	Customer_Id int  `json:"customer_id"` // Required field
	Date        *int `json:"date"`        // Required field

}

type PostDelivery struct {
	Customer_Id int  `json:"customer_id"` // Required field
	Date        *int `json:"date"`
}

type DeliveryUri struct {
	ID int `uri:"id" binding:"required,number"`
}
