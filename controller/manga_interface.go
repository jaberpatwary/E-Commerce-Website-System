package controller

import (
	"github.com/gin-gonic/gin"
)

type MangaControllerInterface interface {
	//category
	InsertCategory(*gin.Context)
	GetAllCategory(*gin.Context)
	GetOneCategory(*gin.Context)
	UpdateCategory(*gin.Context)
	DeleteCategory(*gin.Context)
	//payment
	InsertPayment(*gin.Context)
	GetAllPayment(*gin.Context)
	GetOnePayment(*gin.Context)
	UpdatePayment(*gin.Context)
	DeletePayment(*gin.Context)
	//product
	InsertProduct(*gin.Context)
	GetAllProduct(*gin.Context)
	GetOneProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
	//seller
	InsertSeller(*gin.Context)
	GetAllSeller(*gin.Context)
	GetOneSeller(*gin.Context)
	UpdateSeller(*gin.Context)
	DeleteSeller(*gin.Context)
	//customer
	InsertCustomer(*gin.Context)
	GetAllCustomer(*gin.Context)
	GetOneCustomer(*gin.Context)
	UpdateCustomer(*gin.Context)
	DeleteCustomer(*gin.Context)
	//order
	InsertOrder(*gin.Context)
	GetAllOrder(*gin.Context)
	GetOneOrder(*gin.Context)
	UpdateOrder(*gin.Context)
	DeleteOrder(*gin.Context)
	//delivery
	InsertDelivery(*gin.Context)
	GetAllDelivery(*gin.Context)
	GetOneDelivery(*gin.Context)
	UpdateDelivery(*gin.Context)
	DeleteDelivery(*gin.Context)
	//TransactionReport
	InsertTransactionreport(*gin.Context)
	GetAllTransactionreport(*gin.Context)
	GetOneTransactionreport(*gin.Context)
	UpdateTransactionreport(*gin.Context)
	DeleteTransactionreport(*gin.Context)
}

type CategoryControllerInterface interface {
	//InsertCategory(*gin.Context)
	//GetAllCategory(*gin.Context)
	//GetOneCategory(*gin.Context)
	//UpdateCategory(*gin.Context)
	//DeleteCategory(*gin.Context)
}
type PaymentControllerInterface interface {
	//InsertPayment(*gin.Context)
	//GetAllPayment(*gin.Context)
	//GetOnePayment(*gin.Context)
	//UpdatePayment(*gin.Context)
	//DeletePayment(*gin.Context)
}
type ProductControllerInterface interface {
	//InsertProduct(*gin.Context)
	//GetAllProduct(*gin.Context)
	//GetOneProduct(*gin.Context)
	//UpdateProduct(*gin.Context)
	//DeleteProduct(*gin.Context)
}
type SellerControllerInterface interface {
	//InsertSeller(*gin.Context)
	//GetAllSeller(*gin.Context)
	//GetOneSeller(*gin.Context)
	//UpdateSeller(*gin.Context)
	//DeleteSeller(*gin.Context)
}
type CustomerControllerInterface interface {
	//InsertCustomer(*gin.Context)
	//GetAllCustomer(*gin.Context)
	//GetOneCustomer(*gin.Context)
	//UpdateCustomer(*gin.Context)
	//DeleteCustomer(*gin.Context)
}
type OrderControllerInterface interface {
	//InsertOrder(*gin.Context)
	//GetAllOrder(*gin.Context)
	//GetOneOrder(*gin.Context)
	//UpdateOrder(*gin.Context)
	//DeleteOrder(*gin.Context)
}
type DeliveryControllerInterface interface {
	//InsertDelivery(*gin.Context)
	//GetAllDelivery(*gin.Context)
	//GetOneDelivery(*gin.Context)
	//UpdateDelivery(*gin.Context)
	//DeleteDelivery(*gin.Context)
}
type TransactionreporttControllerInterface interface {
	//InsertTransactionreport(*gin.Context)
	//GetAllTransactionreport(*gin.Context)
	//GetOneTransactionreport(*gin.Context)
	//UpdateTransactionreport(*gin.Context)
	//DeleteTransactionreport(*gin.Context)
}
