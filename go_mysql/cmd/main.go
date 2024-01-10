package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GauravRasal07/go_bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)

	fmt.Println("Starting server at 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}