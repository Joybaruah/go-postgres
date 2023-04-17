package middleware

import (
	"encoding/json"
	"fmt"
	"go-postgres/database"
	"go-postgres/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type responce struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode request body. %v", err)
	}

	insertID := database.InsertStock(stock)

	res := responce{
		ID:      insertID,
		Message: "Stock Created.",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int. %v", err)
	}

	stock, err := database.GetStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := database.GetAllStocks()

	if err != nil {
		log.Fatalf("Unable to get all the stocks. %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int. %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode request body %v", err)
	}

	updatedRows := database.UpdateStock(int64(id), stock)
	msg := fmt.Sprintf("Stock updated. Total records affected %v", updatedRows)

	res := responce{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int. %v", err)
	}

	deletedRows := database.DeleteStock(int64(id))
	msg := fmt.Sprintf("Stock deleted. Total records affected %v", deletedRows)

	res := responce{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
