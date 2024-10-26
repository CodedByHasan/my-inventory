package main

import "database/sql"

type products struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]products, error) {
	query := "SELECT id, name, quantity, price from products"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	product := []products{}

	for rows.Next() {
		var p products

		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)

		if err != nil {
			return nil, err
		}
		product = append(product, p)

	}
	return product, nil
}
