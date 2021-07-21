package app

import (
	"github.com/gorilla/mux"
	"github.com/nicoletafratila/bookstore_items-api/src/clients/elasticsearch"
	"github.com/nicoletafratila/bookstore_utils-go/logger"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapURLs()

	srv := &http.Server{
		Addr:              ":8081",
		WriteTimeout:      50 * time.Second,
		ReadHeaderTimeout: 50 * time.Second,
		IdleTimeout:       50 * time.Second,
		Handler:           router,
	}

	logger.Info("about to start the application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
