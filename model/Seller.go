package model

type Seller struct {
	ID         int     `json:"id"`
	Product_Id int     `json:"product_id"` // Required field
	Name       *string `json:"name"`       // Required field

}

type PostSeller struct {
	Product_Id int     `json:"product_id"` // Required field
	Name       *string `json:"name"`       // Required field
}

type SellerUri struct {
	ID int `uri:"id" binding:"required,number"`
}
