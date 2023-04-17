package database

import (
	"database/sql"
	"fmt"
	"go-postgres/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func createConnection() *sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading env file.")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connnected to DB")
	return db

}

func InsertStock(stock models.Stock) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`

	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute query. %v", err)
	}

	fmt.Printf("Stock Created %v", id)
	return id
}

func GetStock(id int64) (models.Stock, error) {

	db := createConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were Returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to Scan the row %v", err)
	}

	return stock, err
}

func GetAllStocks() ([]models.Stock, error) {

	db := createConnection()
	defer db.Close()

	var stocks []models.Stock
	sqlStatement := `SELECT * FROM stocks`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to Scan the row %v", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, err
}

func UpdateStock(id int64, stock models.Stock) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteStock(id int64) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
