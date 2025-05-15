package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type CustomerRepositoryInterface interface {
	InsertCustomer(model.PostCustomer) bool
	GetAllCustomer() []model.Customer
	GetOneCustomer(int) model.Customer
	DeleteCustomer(int) bool
	UpdateCustomer(int, model.PostCustomer) model.Customer
}
