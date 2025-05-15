package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type OrderRepositoryInterface interface {
	InsertOrder(model.PostOrder) bool
	GetAllOrder() []model.Order
	GetOneOrder(int) model.Order
	DeleteOrder(int) bool
	UpdateOrder(int, model.PostOrder) model.Order
}
