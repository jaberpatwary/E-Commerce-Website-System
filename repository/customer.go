package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type CustomerRepository struct {
	Db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepositoryInterface {
	return &CustomerRepository{Db: db}
}

// Insert customer
func (p *CustomerRepository) InsertCustomer(post model.PostCustomer) bool {
	stmt, err := p.Db.Prepare("INSERT INTO customer(name, address, phone) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Name, post.Address, post.Phone)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll customer
func (p *CustomerRepository) GetAllCustomer() []model.Customer {
	query, err := p.Db.Query("SELECT * FROM eco.customer")
	if err != nil {
		log.Println(err)
		return nil
	}
	var customers []model.Customer
	if query != nil {
		for query.Next() {
			var (
				id      int
				name    string
				address *string
				phone   *int
			)
			err := query.Scan(&id, &name, &address, &phone)
			if err != nil {
				log.Println(err)
			}
			customer := model.Customer{ID: id, Name: name, Address: address, Phone: phone}
			customers = append(customers, customer)
		}
	}
	return customers
}

// GetOne customer
func (m *CustomerRepository) GetOneCustomer(id int) model.Customer {
	query, err := m.Db.Query("SELECT * FROM eco.customer WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Customer{}
	}
	defer query.Close()
	var customer model.Customer
	if query != nil {
		for query.Next() {
			var (
				id      int
				name    string
				address *string
				phone   *int
			)
			err := query.Scan(&id, &name, &address, &phone)
			if err != nil {
				log.Println(err)
				return model.Customer{}
			}
			customer = model.Customer{ID: id, Name: name, Address: address, Phone: phone}
		}
	}
	return customer
}

//update customer

func (p *CustomerRepository) UpdateCustomer(id int, post model.PostCustomer) model.Customer {
	_, err := p.Db.Exec("UPDATE customer SET name = ?, address = ?, phone =? WHERE id = ?", post.Name, post.Address, post.Phone, id)
	if err != nil {
		log.Println(err)
		return model.Customer{}
	}
	return p.GetOneCustomer(id)
}

// Delete customer
func (m *CustomerRepository) DeleteCustomer(id int) bool {
	_, err := m.Db.Exec("DELETE FROM customer WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
