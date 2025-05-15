package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
	"github.com/ramadhanalfarisi/go-simple-crud/repository"
)

type MangaController struct {
	Db *sql.DB
}

func NewMangaController(db *sql.DB) MangaControllerInterface {
	return &MangaController{Db: db}
}

/*InsertPatient

func (m *MangaController) InsertPatient(c *gin.Context) {

	DB := m.Db
	var post model.PostPatient
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewPatientRepository(DB)
	insert := repository.InsertPatient(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " Patient has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " Patient  is not saved!"})
		return
	}

}

// Gell all Patient
func (m *MangaController) GetAllPatient(c *gin.Context) {
	DB := m.Db
	repository := repository.NewPatientRepository(DB)
	get := repository.GetAllPatient()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get Patient successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "Patient not found"})
		return
	}
}

// Get One Patient
func (m *MangaController) GetOnePatient(c *gin.Context) {
	DB := m.Db
	var uri model.PatientUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewPatientRepository(DB)
	get := repository.GetOnePatient(uri.ID)
	if (get != model.Patient{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get patient successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "patient not found"})
		return
	}
}

//update patient

func (m *MangaController) UpdatePatient(c *gin.Context) {
	DB := m.Db
	var post model.PostPatient
	var uri model.PatientUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewPatientRepository(DB)
	update := repository.UpdatePatient(uri.ID, post)
	if (update != model.Patient{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "Patient": "update patient successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "patient": "update patient failed"})
		return
	}
}

//Deletepatient

func (m *MangaController) DeletePatient(c *gin.Context) {
	DB := m.Db
	var uri model.PatientUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewPatientRepository(DB)
	delete := repository.DeletePatient(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete patient successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete patient failed"})
		return
	}
}*/

//Insertcategory

func (m *MangaController) InsertCategory(c *gin.Context) {

	DB := m.Db
	var post model.PostCategory
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewCategoryRepository(DB)
	insert := repository.InsertCategory(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " category has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " category  is not saved!"})
		return
	}

}

// Gell all category
func (m *MangaController) GetAllCategory(c *gin.Context) {
	DB := m.Db
	repository := repository.NewCategoryRepository(DB)
	get := repository.GetAllCategory()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get category successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "category not found"})
		return
	}
}

// Get One cagetory
func (m *MangaController) GetOneCategory(c *gin.Context) {
	DB := m.Db
	var uri model.CategoryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCategoryRepository(DB)
	get := repository.GetOneCategory(uri.ID)
	if (get != model.Category{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get category successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "category not found"})
		return
	}
}

//update category

func (m *MangaController) UpdateCategory(c *gin.Context) {
	DB := m.Db
	var post model.PostCategory
	var uri model.CategoryUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCategoryRepository(DB)
	update := repository.UpdateCategory(uri.ID, post)
	if (update != model.Category{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "category": "update category successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "category": "update category failed"})
		return
	}
}

//Deletecategory

func (m *MangaController) DeleteCategory(c *gin.Context) {
	DB := m.Db
	var uri model.CategoryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCategoryRepository(DB)
	delete := repository.DeleteCategory(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete category successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete category failed"})
		return
	}
}

//Insertpayment

func (m *MangaController) InsertPayment(c *gin.Context) {

	DB := m.Db
	var post model.PostPayment
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewPaymentRepository(DB)
	insert := repository.InsertPayment(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " payment has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " payment  is not saved!"})
		return
	}

}

// Gell all payment
func (m *MangaController) GetAllPayment(c *gin.Context) {
	DB := m.Db
	repository := repository.NewPaymentRepository(DB)
	get := repository.GetAllPayment()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get payment successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "payment not found"})
		return
	}
}

// Get One payment
func (m *MangaController) GetOnePayment(c *gin.Context) {
	DB := m.Db
	var uri model.PaymentUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewPaymentRepository(DB)
	get := repository.GetOnePayment(uri.ID)
	if (get != model.Payment{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get payment successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": " payment not found"})
		return
	}
}

//update payment

func (m *MangaController) UpdatePayment(c *gin.Context) {
	DB := m.Db
	var post model.PostPayment
	var uri model.PaymentUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewPaymentRepository(DB)
	update := repository.UpdatePayment(uri.ID, post)
	if (update != model.Payment{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "payment": "update payment successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "payment": "update payment failed"})
		return
	}
}

//Deletepayment

func (m *MangaController) DeletePayment(c *gin.Context) {
	DB := m.Db
	var uri model.PaymentUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewPaymentRepository(DB)
	delete := repository.DeletePayment(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete payment successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete payment failed"})
		return
	}
}

//Insertproduct

func (m *MangaController) InsertProduct(c *gin.Context) {

	DB := m.Db
	var post model.PostProduct
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewProductRepository(DB)
	insert := repository.InsertProduct(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " product has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " product  is not saved!"})
		return
	}

}

// Gell all product
func (m *MangaController) GetAllProduct(c *gin.Context) {
	DB := m.Db
	repository := repository.NewProductRepository(DB)
	get := repository.GetAllProduct()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get product successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "product not found"})
		return
	}
}

// Get One product
func (m *MangaController) GetOneProduct(c *gin.Context) {
	DB := m.Db
	var uri model.ProductUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewProductRepository(DB)
	get := repository.GetOneProduct(uri.ID)
	if (get != model.Product{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get product successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": " product not found"})
		return
	}
}

//update product

func (m *MangaController) UpdateProduct(c *gin.Context) {
	DB := m.Db
	var post model.PostProduct
	var uri model.ProductUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewProductRepository(DB)
	update := repository.UpdateProduct(uri.ID, post)
	if (update != model.Product{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "product": "update product successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "product": "update product failed"})
		return
	}
}

//Deleteproduct

func (m *MangaController) DeleteProduct(c *gin.Context) {
	DB := m.Db
	var uri model.PaymentUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewProductRepository(DB)
	delete := repository.DeleteProduct(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete product successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete product failed"})
		return
	}
}

//Insertseller

func (m *MangaController) InsertSeller(c *gin.Context) {

	DB := m.Db
	var post model.PostSeller
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewSellerRepository(DB)
	insert := repository.InsertSeller(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " seller has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " seller  is not saved!"})
		return
	}

}

// Gell all seller
func (m *MangaController) GetAllSeller(c *gin.Context) {
	DB := m.Db
	repository := repository.NewSellerRepository(DB)
	get := repository.GetAllSeller()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get seller successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "seller not found"})
		return
	}
}

// Get One seller
func (m *MangaController) GetOneSeller(c *gin.Context) {
	DB := m.Db
	var uri model.SellerUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewSellerRepository(DB)
	get := repository.GetOneSeller(uri.ID)
	if (get != model.Seller{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get seller successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": " seller not found"})
		return
	}
}

//update seller

func (m *MangaController) UpdateSeller(c *gin.Context) {
	DB := m.Db
	var post model.PostSeller
	var uri model.SellerUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewSellerRepository(DB)
	update := repository.UpdateSeller(uri.ID, post)
	if (update != model.Seller{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "seller": "update seller successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "seller": "update seller failed"})
		return
	}
}

//Deleteseller

func (m *MangaController) DeleteSeller(c *gin.Context) {
	DB := m.Db
	var uri model.SellerUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewSellerRepository(DB)
	delete := repository.DeleteSeller(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete seller successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete seller failed"})
		return
	}
}

//Insertcustomer

func (m *MangaController) InsertCustomer(c *gin.Context) {

	DB := m.Db
	var post model.PostCustomer
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewCustomerRepository(DB)
	insert := repository.InsertCustomer(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " customer has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " customer  is not saved!"})
		return
	}

}

// Gell all customer
func (m *MangaController) GetAllCustomer(c *gin.Context) {
	DB := m.Db
	repository := repository.NewCustomerRepository(DB)
	get := repository.GetAllCustomer()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get customer successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "customer not found"})
		return
	}
}

// Get One customer
func (m *MangaController) GetOneCustomer(c *gin.Context) {
	DB := m.Db
	var uri model.CustomerUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCustomerRepository(DB)
	get := repository.GetOneCustomer(uri.ID)
	if (get != model.Customer{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get customer successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": " customer not found"})
		return
	}
}

//update customer

func (m *MangaController) UpdateCustomer(c *gin.Context) {
	DB := m.Db
	var post model.PostCustomer
	var uri model.CustomerUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCustomerRepository(DB)
	update := repository.UpdateCustomer(uri.ID, post)
	if (update != model.Customer{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "customer": "update customer successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "customer": "update customer failed"})
		return
	}
}

//Deletecustomer

func (m *MangaController) DeleteCustomer(c *gin.Context) {
	DB := m.Db
	var uri model.CustomerUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewCustomerRepository(DB)
	delete := repository.DeleteCustomer(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete customer successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete customre failed"})
		return
	}
}

//Insertorder

func (m *MangaController) InsertOrder(c *gin.Context) {

	DB := m.Db
	var post model.PostOrder
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewOrderRepository(DB)
	insert := repository.InsertOrder(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " order has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " order  is not saved!"})
		return
	}

}

// Gell all order
func (m *MangaController) GetAllOrder(c *gin.Context) {
	DB := m.Db
	repository := repository.NewOrderRepository(DB)
	get := repository.GetAllOrder()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get order successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "order not found"})
		return
	}
}

// Get One order
func (m *MangaController) GetOneOrder(c *gin.Context) {
	DB := m.Db
	var uri model.OrderUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewOrderRepository(DB)
	get := repository.GetOneOrder(uri.ID)
	if (get != model.Order{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get order successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": " order not found"})
		return
	}
}

//update order

func (m *MangaController) UpdateOrder(c *gin.Context) {
	DB := m.Db
	var post model.PostOrder
	var uri model.OrderUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewOrderRepository(DB)
	update := repository.UpdateOrder(uri.ID, post)
	if (update != model.Order{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "order": "update order successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "order": "update order failed"})
		return
	}
}

//Delete order

func (m *MangaController) DeleteOrder(c *gin.Context) {
	DB := m.Db
	var uri model.OrderUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewOrderRepository(DB)
	delete := repository.DeleteOrder(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete order successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete order failed"})
		return
	}
}

//Insertdelivery

func (m *MangaController) InsertDelivery(c *gin.Context) {

	DB := m.Db
	var post model.PostDelivery
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewDeliveryRepository(DB)
	insert := repository.InsertDelivery(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " delivery has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": "delivery  is not saved!"})
		return
	}

}

// Gell all delivery
func (m *MangaController) GetAllDelivery(c *gin.Context) {
	DB := m.Db
	repository := repository.NewDeliveryRepository(DB)
	get := repository.GetAllDelivery()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get delivery successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "delivery not found"})
		return
	}
}

// Get One delivery
func (m *MangaController) GetOneDelivery(c *gin.Context) {
	DB := m.Db
	var uri model.DeliveryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewDeliveryRepository(DB)
	get := repository.GetOneDelivery(uri.ID)
	if (get != model.Delivery{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get delivery successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": " delivery not found"})
		return
	}
}

//update delivery

func (m *MangaController) UpdateDelivery(c *gin.Context) {
	DB := m.Db
	var post model.PostDelivery
	var uri model.DeliveryUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewDeliveryRepository(DB)
	update := repository.UpdateDelivery(uri.ID, post)
	if (update != model.Delivery{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "delivery": "update delivery successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "delivery": "update deliveryfailed"})
		return
	}
}

//Delete delivery

func (m *MangaController) DeleteDelivery(c *gin.Context) {
	DB := m.Db
	var uri model.DeliveryUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewDeliveryRepository(DB)
	delete := repository.DeleteDelivery(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete delivery successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete delivery failed"})
		return
	}
}

//Inserttansactionreport

func (m *MangaController) InsertTransactionreport(c *gin.Context) {

	DB := m.Db
	var post model.PostTransactionreport
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewTransactionreportRepository(DB)
	insert := repository.InsertTransactionreport(post)
	if insert {
		c.JSON(201, gin.H{"Status": "success", "meg": " Transactionreport has saved!"})
		return
	} else {
		c.JSON(500, gin.H{"Status": "failed", "meg": " Transactionreportt  is not saved!"})
		return
	}

}

// Gell all tansantionrepot
func (m *MangaController) GetAllTransactionreport(c *gin.Context) {
	DB := m.Db
	repository := repository.NewTransactionreportRepository(DB)
	get := repository.GetAllTransactionreport()
	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get tansactionrepot successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": "tansactionreport not found"})
		return
	}
}

// Get One tansactionreport
func (m *MangaController) GetOneTransactionreport(c *gin.Context) {
	DB := m.Db
	var uri model.TransactionreportUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTransactionreportRepository(DB)
	get := repository.GetOneTransactionreport(uri.ID)
	if (get != model.Transactionreport{}) {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get transactionreport successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "success", "data": nil, "msg": " transactionreport not found"})
		return
	}
}

//update transactionreport

func (m *MangaController) UpdateTransactionreport(c *gin.Context) {
	DB := m.Db
	var post model.PostTransactionreport
	var uri model.TransactionreportUri
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTransactionreportRepository(DB)
	update := repository.UpdateTransactionreport(uri.ID, post)
	if (update != model.Transactionreport{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "transactionreport": "update transactionreport successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "transactionreport": "update transactionreport failed"})
		return
	}
}

//Delete transactionreport

func (m *MangaController) DeleteTransactionreport(c *gin.Context) {
	DB := m.Db
	var uri model.TransactionreportUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}
	repository := repository.NewTransactionreportRepository(DB)
	delete := repository.DeleteTransactionreport(uri.ID)
	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete transactionreport successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete transactinreport failed"})
		return
	}
}
