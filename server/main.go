package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"./router"
)

func main() {
	r := router.Router()

	n := negroni.Classic()
	n.UseHandler(r)

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", n))
}
