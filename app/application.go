package app

import (
	"github.com/gorilla/mux"
	"github.com/nicoletafratila/bookstore_items-api/clients/elasticsearch"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapURLs()

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
