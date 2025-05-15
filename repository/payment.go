package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type PaymentRepository struct {
	Db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepositoryInterface {
	return &PaymentRepository{Db: db}
}

// Insert payment
func (p *PaymentRepository) InsertPayment(post model.PostPayment) bool {
	stmt, err := p.Db.Prepare("INSERT INTO payment(category_id, date) VALUES (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Category_Id, post.Date)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll payment
func (p *PaymentRepository) GetAllPayment() []model.Payment {
	query, err := p.Db.Query("SELECT * FROM eco.payment")
	if err != nil {
		log.Println(err)
		return nil
	}
	var payments []model.Payment
	if query != nil {
		for query.Next() {
			var (
				id          int
				category_id int
				date        *int
			)
			err := query.Scan(&id, &category_id, &date)
			if err != nil {
				log.Println(err)
			}
			Payment := model.Payment{ID: id, Category_Id: category_id, Date: date}
			payments = append(payments, Payment)
		}
	}
	return payments
}

// GetOne Payment
func (m *PaymentRepository) GetOnePayment(id int) model.Payment {
	query, err := m.Db.Query("SELECT * FROM eco.payment WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Payment{}
	}
	defer query.Close()
	var payment model.Payment
	if query != nil {
		for query.Next() {
			var (
				id          int
				category_id int
				date        *int
			)
			err := query.Scan(&id, &category_id, &date)
			if err != nil {
				log.Println(err)
				return model.Payment{}
			}
			payment = model.Payment{ID: id, Category_Id: category_id, Date: date}
		}
	}
	return payment
}

//update payment

func (p *PaymentRepository) UpdatePayment(id int, post model.PostPayment) model.Payment {
	_, err := p.Db.Exec("UPDATE payment SET category_id = ?, date = ? WHERE id = ?", post.Category_Id, post.Date, id)
	if err != nil {
		log.Println(err)
		return model.Payment{}
	}
	return p.GetOnePayment(id)
}

// Delete payment
func (m *PaymentRepository) DeletePayment(id int) bool {
	_, err := m.Db.Exec("DELETE FROM payment WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
