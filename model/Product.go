package model

type Product struct {
	ID          int     `json:"id"`
	Category_Id int     `json:"category_id"` // Required field
	Name        *string `json:"name"`        // Required field

}

type PostProduct struct {
	Category_Id int     `json:"category_id"` // Required field
	Name        *string `json:"name"`
}

type ProductUri struct {
	ID int `uri:"id" binding:"required,number"`
}
