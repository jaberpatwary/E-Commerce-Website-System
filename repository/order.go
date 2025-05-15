package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepositoryInterface {
	return &OrderRepository{Db: db}
}

// Insert order
func (p *OrderRepository) InsertOrder(post model.PostOrder) bool {
	stmt, err := p.Db.Prepare("INSERT INTO `order`(customer_id, date) VALUES (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Customer_Id, post.Date)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll oder
func (p *OrderRepository) GetAllOrder() []model.Order {
	query, err := p.Db.Query("SELECT * FROM eco.`order`")
	if err != nil {
		log.Println(err)
		return nil
	}
	var orders []model.Order
	if query != nil {
		for query.Next() {
			var (
				id          int
				customer_id int
				date        *int
			)
			err := query.Scan(&id, &customer_id, &date)
			if err != nil {
				log.Println(err)
			}
			Order := model.Order{ID: id, Customer_Id: customer_id, Date: date}
			orders = append(orders, Order)
		}
	}
	return orders
}

// GetOne oder
func (m *OrderRepository) GetOneOrder(id int) model.Order {
	query, err := m.Db.Query("SELECT * FROM eco.`order` WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Order{}
	}
	defer query.Close()
	var order model.Order
	if query != nil {
		for query.Next() {
			var (
				id          int
				customer_id int
				date        *int
			)
			err := query.Scan(&id, &customer_id, &date)
			if err != nil {
				log.Println(err)
				return model.Order{}
			}
			order = model.Order{ID: id, Customer_Id: customer_id, Date: date}
		}
	}
	return order
}

//update order

func (p *OrderRepository) UpdateOrder(id int, post model.PostOrder) model.Order {
	_, err := p.Db.Exec("UPDATE `order` SET customer_id = ?, date = ? WHERE id = ?", post.Customer_Id, post.Date, id)
	if err != nil {
		log.Println(err)
		return model.Order{}
	}
	return p.GetOneOrder(id)
}

// Delete order
func (m *OrderRepository) DeleteOrder(id int) bool {
	_, err := m.Db.Exec("DELETE FROM order WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
