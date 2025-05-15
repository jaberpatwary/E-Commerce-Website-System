package repository

import "github.com/ramadhanalfarisi/go-simple-crud/model"

type DeliveryRepositoryInterface interface {
	InsertDelivery(model.PostDelivery) bool
	GetAllDelivery() []model.Delivery
	GetOneDelivery(int) model.Delivery
	DeleteDelivery(int) bool
	UpdateDelivery(int, model.PostDelivery) model.Delivery
}
