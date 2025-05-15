package model

type Customer struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`    // Required field
	Address *string `json:"address"` // Required field
	Phone   *int    `json:"phone"`
}

type PostCustomer struct {
	Name    string  `json:"name"`    // Required field
	Address *string `json:"address"` // Required field
	Phone   *int    `json:"phone"`
}

type CustomerUri struct {
	ID int `uri:"id" binding:"required,number"`
}
