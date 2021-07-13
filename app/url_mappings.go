package app

import (
	"github.com/nicoletafratila/bookstore_items-api/controllers"
	"net/http"
)

func mapURLs() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
