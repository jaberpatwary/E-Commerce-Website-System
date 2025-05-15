package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type ProductRepositoryInterface interface {
	InsertProduct(model.PostProduct) bool
	GetAllProduct() []model.Product
	GetOneProduct(int) model.Product
	DeleteProduct(int) bool
	UpdateProduct(int, model.PostProduct) model.Product
}
