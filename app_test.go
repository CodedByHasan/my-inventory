package main

import (
	"fmt"
	"log"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	db_user := loadEnvVar("DB_USER")
	db_password := loadEnvVar("DB_PASSWORD")

	err := a.Initialise(db_user, db_password, "test")
	if err != nil {
		log.Fatal("Error occured while initialising the database")
	}

	createTable()
	m.Run()
}

func createTable() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS products (
	id int NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	quantity int,
	price float(10,7),
	PRIMARY KEY (id)
	);`

	_, err := a.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM PRODUCTS")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT into products(name, quantity, price) VALUES('%v, %v, %v')", name, quantity, price)
	a.DB.Exec(query)
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("keyboard", 100, 125)
}
