package model

import (
	"database/sql"
	"errors"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func GetProducts(db *sql.DB) ([]Product, error) {
	query := "SELECT id, name, quantity, price from products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (p *Product) GetProduct(db *sql.DB) error {
	query := fmt.Sprintf("SELECT name, quantity, price FROM products where id=%v", p.ID)
	row := db.QueryRow(query)

	err := row.Scan(&p.Name, &p.Quantity, &p.Price)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) CreateProduct(db *sql.DB) error {
	query := fmt.Sprintf("insert into products(name,quantity,price) values('%v', %v, %v)", p.Name, p.Quantity, p.Price)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = int(id)
	return nil
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	query := fmt.Sprintf("update products set name='%v', quantity=%v, price=%v where id=%v", p.Name, p.Quantity, p.Price, p.ID)
	result, _ := db.Exec(query)

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no such row exists")
	}
	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	query := fmt.Sprintf("delete from products where id=%v", p.ID)

	_, err := db.Exec(query)
	return err
}
