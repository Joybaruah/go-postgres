package main

import(
	"log"
	"fmt"
	"net/http"
	"go-postgres/router"
)

func main() {
	r := router.Router()
	fmt.Println("Running on 8080..")

	log.Fatal(http.ListenAndServe(":8080", r))
}