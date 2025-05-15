package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type TransactionreportRepository struct {
	Db *sql.DB
}

func NewTransactionreportRepository(db *sql.DB) TransactionreportRepositoryInterface {
	return &TransactionreportRepository{Db: db}
}

// Insert Transactionreport
func (p *TransactionreportRepository) InsertTransactionreport(post model.PostTransactionreport) bool {
	stmt, err := p.Db.Prepare("INSERT INTO transactionreport(customer_id, order_id, product_id, payment_id) VALUES (?,?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Customer_Id, post.Order_Id, post.Product_Id, post.Payment_Id)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll tansactionreport
func (p *TransactionreportRepository) GetAllTransactionreport() []model.Transactionreport {
	query, err := p.Db.Query("SELECT * FROM eco.transactionreport")
	if err != nil {
		log.Println(err)
		return nil
	}
	var transactionreports []model.Transactionreport
	if query != nil {
		for query.Next() {
			var (
				id          int
				customer_id int
				order_id    *int
				product_id  *int
				payment_id  *int
			)
			err := query.Scan(&id, &customer_id, &order_id, &product_id, &payment_id)
			if err != nil {
				log.Println(err)
			}
			TransactionReport := model.Transactionreport{ID: id, Customer_Id: customer_id, Order_Id: order_id, Product_Id: product_id, Payment_Id: payment_id}
			transactionreports = append(transactionreports, TransactionReport)
		}
	}
	return transactionreports
}

// GetOne transactionreport
func (m *TransactionreportRepository) GetOneTransactionreport(id int) model.Transactionreport {
	query, err := m.Db.Query("SELECT * FROM eco.transactionreport WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Transactionreport{}
	}
	defer query.Close()
	var tansactionreport model.Transactionreport
	if query != nil {
		for query.Next() {
			var (
				id          int
				customer_id int
				order_id    *int
				product_id  *int
				payment_id  *int
			)
			err := query.Scan(&id, &customer_id, &order_id, &product_id, &payment_id)
			if err != nil {
				log.Println(err)
				return model.Transactionreport{}
			}
			tansactionreport = model.Transactionreport{ID: id, Customer_Id: customer_id, Order_Id: order_id, Product_Id: product_id, Payment_Id: payment_id}
		}
	}
	return tansactionreport
}

//update tansactionreport

func (p *TransactionreportRepository) UpdateTransactionreport(id int, post model.PostTransactionreport) model.Transactionreport {
	_, err := p.Db.Exec("UPDATE transactionreport SET customer_id = ?, order_id = ?, product_id =?, payment_id =? WHERE id = ?", post.Customer_Id, post.Order_Id, post.Product_Id, post.Payment_Id, id)
	if err != nil {
		log.Println(err)
		return model.Transactionreport{}
	}
	return p.GetOneTransactionreport(id)
}

// Delete transactionreport
func (m *TransactionreportRepository) DeleteTransactionreport(id int) bool {
	_, err := m.Db.Exec("DELETE FROM transactionreport WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
