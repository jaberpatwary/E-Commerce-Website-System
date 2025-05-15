package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type DeliveryRepository struct {
	Db *sql.DB
}

func NewDeliveryRepository(db *sql.DB) DeliveryRepositoryInterface {
	return &DeliveryRepository{Db: db}
}

// Insert delivery
func (p *DeliveryRepository) InsertDelivery(post model.PostDelivery) bool {
	stmt, err := p.Db.Prepare("INSERT INTO delivery(customer_id, date) VALUES (?,?)")
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

// GetAll delivery
func (p *DeliveryRepository) GetAllDelivery() []model.Delivery {
	query, err := p.Db.Query("SELECT * FROM eco.delivery")
	if err != nil {
		log.Println(err)
		return nil
	}
	var deliverys []model.Delivery
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
			Delivery := model.Delivery{ID: id, Customer_Id: customer_id, Date: date}
			deliverys = append(deliverys, Delivery)
		}
	}
	return deliverys
}

// GetOne delivery
func (m *DeliveryRepository) GetOneDelivery(id int) model.Delivery {
	query, err := m.Db.Query("SELECT * FROM eco.delivery WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Delivery{}
	}
	defer query.Close()
	var delivery model.Delivery
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
				return model.Delivery{}
			}
			delivery = model.Delivery{ID: id, Customer_Id: customer_id, Date: date}
		}
	}
	return delivery
}

//update delivery

func (p *DeliveryRepository) UpdateDelivery(id int, post model.PostDelivery) model.Delivery {
	_, err := p.Db.Exec("UPDATE delivery SET customer_id = ?, date = ? WHERE id = ?", post.Customer_Id, post.Date, id)
	if err != nil {
		log.Println(err)
		return model.Delivery{}
	}
	return p.GetOneDelivery(id)
}

// Delete delivery
func (m *DeliveryRepository) DeleteDelivery(id int) bool {
	_, err := m.Db.Exec("DELETE FROM delivery WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
