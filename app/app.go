package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-simple-crud/controller"
	"github.com/ramadhanalfarisi/go-simple-crud/db"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConnection() {
	db := db.Connectdb()
	a.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewMangaController(a.DB)

	/*patient router
	r.POST("/patient", controller.InsertPatient)
	r.GET("/patient", controller.GetAllPatient)
	r.GET("/patient/:id", controller.GetOnePatient)
	r.PUT("/patient/:id", controller.UpdatePatient)
	r.DELETE("/patient/:id", controller.DeletePatient)*/

	//category router
	r.POST("/category", controller.InsertCategory)
	r.GET("/category", controller.GetAllCategory)
	r.GET("/category/:id", controller.GetOneCategory)
	r.PUT("/category/:id", controller.UpdateCategory)
	r.DELETE("/category/:id", controller.DeleteCategory)
	//payment router
	r.POST("/payment", controller.InsertPayment)
	r.GET("/payment", controller.GetAllPayment)
	r.GET("/payment/:id", controller.GetOnePayment)
	r.PUT("/payment/:id", controller.UpdatePayment)
	r.DELETE("/payment/:id", controller.DeletePayment)
	//product router
	r.POST("/product", controller.InsertProduct)
	r.GET("/product", controller.GetAllProduct)
	r.GET("/product/:id", controller.GetOneProduct)
	r.PUT("/product/:id", controller.UpdateProduct)
	r.DELETE("/product/:id", controller.DeleteProduct)
	//seller router
	r.POST("/seller", controller.InsertSeller)
	r.GET("/seller", controller.GetAllSeller)
	r.GET("/seller/:id", controller.GetOneSeller)
	r.PUT("/seller/:id", controller.UpdateSeller)
	r.DELETE("/seller/:id", controller.DeleteSeller)
	//customer router
	r.POST("/customer", controller.InsertCustomer)
	r.GET("/customer", controller.GetAllCustomer)
	r.GET("/customer/:id", controller.GetOneCustomer)
	r.PUT("/customer/:id", controller.UpdateCustomer)
	r.DELETE("/customer/:id", controller.DeleteCustomer)
	//order router
	r.POST("/order", controller.InsertOrder)
	r.GET("/order", controller.GetAllOrder)
	r.GET("/order/:id", controller.GetOneOrder)
	r.PUT("/order/:id", controller.UpdateOrder)
	r.DELETE("/order/:id", controller.DeleteOrder)
	//delivery router
	r.POST("/delivery", controller.InsertDelivery)
	r.GET("/delivery", controller.GetAllDelivery)
	r.GET("/delivery/:id", controller.GetOneDelivery)
	r.PUT("/delivery/:id", controller.UpdateDelivery)
	r.DELETE("/delivery/:id", controller.DeleteDelivery)
	//tansection report
	r.POST("/transactionreport", controller.InsertTransactionreport)
	r.GET("/transactionreport", controller.GetAllTransactionreport)
	r.GET("/transactionreport/:id", controller.GetOneTransactionreport)
	r.PUT("/transactionreport/:id", controller.UpdateTransactionreport)
	r.DELETE("/transactionreport/:id", controller.DeleteTransactionreport)
	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":1000")
}
