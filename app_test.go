package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
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
	a.DB.Exec("DELETE from products")
	// Every time table is cleared, begin auto increment at 1
	a.DB.Exec("ALTER table products AUTO_INCREMENT=1")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT into products(name, quantity, price) VALUES('%v', %v, %v)", name, quantity, price)
	_, err := a.DB.Exec(query)
	if err != nil {
		log.Println(err)
	}
}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received: %v", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("keyboard", 100, 125)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)
}

func TestCreateProduct(t *testing.T) {
	clearTable()
	var product = []byte(`{"name":"chair", "quantity":1, "price":100}`)
	request, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	request.Header.Set("Content-Type", "application/json")

	response := sendRequest(request)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "chair" {
		t.Errorf("Expected name: %v, Got: %v", "chair", m["name"])
	}

	if m["quantity"] != 1.0 {
		t.Errorf("Expected quantity: %v, Got: %v", 1.0, m["quantity"])
	}
}

func TestDeleteProduct(t *testing.T) {
	clearTable()
	addProduct("test product", 10, 10)

	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("DELETE", "/product/1", nil)
	response = sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

	request, _ = http.NewRequest("GET", "/product/1", nil)
	response = sendRequest(request)
	checkStatusCode(t, http.StatusNotFound, response.Code)
}

func TestUpdateProduct(t *testing.T) {
	clearTable()
	addProduct("test product", 10, 10)

	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

	var oldValue map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &oldValue)

	// changing quantity from 10 -> 1 and price from 10 -> 15
	var product = []byte(`{"name":"test product", "quantity":1, "price":15}`)
	request, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(product))
	request.Header.Set("Content-Type", "application/json")

	response = sendRequest(request)
	var newValue map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &newValue)

	if oldValue["id"] != newValue["id"] {
		t.Errorf("Expected id: %v, Got: %v", newValue["id"], oldValue["id"])
	}

	if oldValue["name"] != newValue["name"] {
		t.Errorf("Expected id: %v, Got: %v", newValue["name"], oldValue["name"])
	}

	if oldValue["quantity"] == newValue["quantity"] {
		t.Errorf("Expected id: %v, Got: %v", newValue["quantity"], oldValue["quantity"])
	}

	if oldValue["price"] == newValue["price"] {
		t.Errorf("Expected id: %v, Got: %v", newValue["price"], oldValue["price"])
	}
}
