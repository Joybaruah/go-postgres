package middleware

import(
	"log"
	"fmt"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
)

func createConnection() &sql.DB{

	err := godotenv.Load(.env)

	if err != nil{
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