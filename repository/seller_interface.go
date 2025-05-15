package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type SellerRepositoryInterface interface {
	InsertSeller(model.PostSeller) bool
	GetAllSeller() []model.Seller
	GetOneSeller(int) model.Seller
	DeleteSeller(int) bool
	UpdateSeller(int, model.PostSeller) model.Seller
}
