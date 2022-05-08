package main

import (
	"bookstore-golang/pkg/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Startin Server at port 9010\n")
	log.Fatal(http.ListenAndServe(":9010", r))

}
