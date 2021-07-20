package main

import (
	"github.com/nicoletafratila/bookstore_items-api/app"
	"os"
)

func main() {
	os.Setenv("LOG_LEVEL", "error")
	app.StartApplication()
}
