package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type SellerRepository struct {
	Db *sql.DB
}

func NewSellerRepository(db *sql.DB) SellerRepositoryInterface {
	return &SellerRepository{Db: db}
}

// Insert seller
func (p *SellerRepository) InsertSeller(post model.PostSeller) bool {
	stmt, err := p.Db.Prepare("INSERT INTO seller(product_id, name) VALUES (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Product_Id, post.Name)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll seller
func (p *SellerRepository) GetAllSeller() []model.Seller {
	query, err := p.Db.Query("SELECT * FROM eco.seller")
	if err != nil {
		log.Println(err)
		return nil
	}
	var sellers []model.Seller
	if query != nil {
		for query.Next() {
			var (
				id         int
				product_id int
				name       *string
			)
			err := query.Scan(&id, &product_id, &name)
			if err != nil {
				log.Println(err)
			}
			Seller := model.Seller{ID: id, Product_Id: product_id, Name: name}
			sellers = append(sellers, Seller)
		}
	}
	return sellers
}

// GetOne seller
func (m *SellerRepository) GetOneSeller(id int) model.Seller {
	query, err := m.Db.Query("SELECT * FROM eco.seller WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Seller{}
	}
	defer query.Close()
	var seller model.Seller
	if query != nil {
		for query.Next() {
			var (
				id         int
				product_id int
				name       *string
			)
			err := query.Scan(&id, &product_id, &name)
			if err != nil {
				log.Println(err)
				return model.Seller{}
			}
			seller = model.Seller{ID: id, Product_Id: product_id, Name: name}
		}
	}
	return seller
}

//update seller

func (p *SellerRepository) UpdateSeller(id int, post model.PostSeller) model.Seller {
	_, err := p.Db.Exec("UPDATE seller SET product_id = ?, name  = ? WHERE id = ?", post.Product_Id, post.Name, id)
	if err != nil {
		log.Println(err)
		return model.Seller{}
	}
	return p.GetOneSeller(id)
}

// Delete seller
func (m *SellerRepository) DeleteSeller(id int) bool {
	_, err := m.Db.Exec("DELETE FROM seller WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
