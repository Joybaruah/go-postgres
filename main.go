package main

import (
	"log"
	"fmt"
	"net/http"
	"go-postgres/router"
)

func main() {
	r := router.Router()
	fmt.PrintIn("Running on 8080..")

	log.fatal(http.ListenAndServe(":8080", r))
}