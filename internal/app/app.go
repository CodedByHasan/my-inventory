package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"my-inventory/internal/model"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func checkDatabaseConnection(database *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	status := "database is up"
	if err := database.PingContext(ctx); err != nil {
		status = "database is down"
	}
	log.Println(status)
}

func (app *App) Initialise(DbUser string, DbPassword string, DbHost string, DbPort string, DbName string) error {
	connectionString := fmt.Sprintf(
		// connection string should be something like: user:password@tcp(host:port)/db-name
		"%s:%s@tcp(%s:%s)/%s",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName,
	)

	fmt.Println(connectionString)

	var err error

	// if DbHost == "localhost" {
	// 	fmt.Println(connectionString)
	// 	} else {
	// 		fmt.Println(connectionString)
	// 	}
	// 	app.DB, err = sql.Open(fmt.Sprint("%s", connectionString), DbHost)
	app.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	checkDatabaseConnection(app.DB)

	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()

	return nil
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := model.GetProducts(app.DB)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)
}

func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	p := model.Product{ID: key}

	err = p.GetProduct(app.DB)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			sendError(w, http.StatusNotFound, "Product not found")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendResponse(w, http.StatusOK, p)
}

func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p model.Product

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = p.CreateProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusCreated, p)
}

func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	var p model.Product
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	p.ID = key
	err = p.UpdateProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, p)
}

func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid product ID")
		return
	}

	p := model.Product{ID: key}
	err = p.DeleteProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"result": "successful deletion"})
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	errorMessage := map[string]string{"error": err}
	sendResponse(w, statusCode, errorMessage)
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/api/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/api/product/{id}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/api/product", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/api/product/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/api/product/{id}", app.deleteProduct).Methods("DELETE")
}
