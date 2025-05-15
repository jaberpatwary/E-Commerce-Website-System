package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type PaymentRepositoryInterface interface {
	InsertPayment(model.PostPayment) bool
	GetAllPayment() []model.Payment
	GetOnePayment(int) model.Payment
	DeletePayment(int) bool
	UpdatePayment(int, model.PostPayment) model.Payment
}
