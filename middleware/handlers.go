package middleware

import(
	"log"
	"fmt"
	"os"
	"net/http"
	"database/sql"
	"go-postgres/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type responce struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() &sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading env file.")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.PrintIn("Connnected to DB")
	return db

}

func CreateStock(w http.ResponceWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Unable to decode request body. %v", err)
	}

	insertID := insertStock(stock)

	res := responce{
		ID: insertID,
		Message: "Stock Created."
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponceWriter, r *http.Request) {
	params = mux.Vars(r)

	id, err = strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int. %v", err)
	}

	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponceWriter, r *http.Request) {
	stocks, err := getAllStocks()

	if err != nil {
		log.Fatalf("Unable to get all the stocks. %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponceWriter, r *http.Request) {
	params = mux.Vars(r)

	id, err = strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int. %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(stock)
	if err != nil {
		log.Fatalf("Unable to decode request body %v", err)
	}

	updatedRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("Stock updated. Total records affected %v", updatedRows)

	res := responce {
		ID: int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponceWriter, r *http.Request) {
	params = mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int. %v", err)
	}

	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("Stock deleted. Total records affected %v", deletedRows)

	res := responce {
		ID: int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// Database functions

func insertStock(stock models.Stock) int64 {

	db := createConnection()
	defer.Close()

	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`

	var id int64
	err := db.QueryRow(sqlStatement, stock.name, stock.price, stock.company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute query. %v", err)
	}

	fmt.PrintIn("Stock Created %v", id)
	return id
}

func getStock(id int64) (models.Stock, error) {
	
	db := createConnection()
	defer.Close()
	
	var stock models.Stock
	
	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.PrintIn("No rows were Returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to Scan the row %v", err)
	}

	return stock, err
}