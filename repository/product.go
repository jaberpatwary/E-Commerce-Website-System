package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramadhanalfarisi/go-simple-crud/model"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryInterface {
	return &ProductRepository{Db: db}
}

// Insert product
func (p *ProductRepository) InsertProduct(post model.PostProduct) bool {
	stmt, err := p.Db.Prepare("INSERT INTO product(category_id, name) VALUES (?,?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Category_Id, post.Name)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// GetAll product
func (p *ProductRepository) GetAllProduct() []model.Product {
	query, err := p.Db.Query("SELECT * FROM eco.product")
	if err != nil {
		log.Println(err)
		return nil
	}
	var products []model.Product
	if query != nil {
		for query.Next() {
			var (
				id          int
				category_id int
				name        *string
			)
			err := query.Scan(&id, &category_id, &name)
			if err != nil {
				log.Println(err)
			}
			Product := model.Product{ID: id, Category_Id: category_id, Name: name}
			products = append(products, Product)
		}
	}
	return products
}

// GetOne Product
func (m *ProductRepository) GetOneProduct(id int) model.Product {
	query, err := m.Db.Query("SELECT * FROM eco.product WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return model.Product{}
	}
	defer query.Close()
	var product model.Product
	if query != nil {
		for query.Next() {
			var (
				id          int
				category_id int
				name        *string
			)
			err := query.Scan(&id, &category_id, &name)
			if err != nil {
				log.Println(err)
				return model.Product{}
			}
			product = model.Product{ID: id, Category_Id: category_id, Name: name}
		}
	}
	return product
}

//update product

func (p *ProductRepository) UpdateProduct(id int, post model.PostProduct) model.Product {
	_, err := p.Db.Exec("UPDATE product SET category_id = ?, name  = ? WHERE id = ?", post.Category_Id, post.Name, id)
	if err != nil {
		log.Println(err)
		return model.Product{}
	}
	return p.GetOneProduct(id)
}

// Delete product
func (m *ProductRepository) DeleteProduct(id int) bool {
	_, err := m.Db.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
