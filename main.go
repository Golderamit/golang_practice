package main

import (
	handler "github/practice22/handlers"
	"github/practice22/storage/postgres"
	"log"
	"net/http"
	"time"
)
func main (){

	dbString := newDBFromConfig()
	store, err := postgres.NewStorage(dbString)
	if err != nil {
		log.Fatal(err)
	   }

	r ,err := handler.NewServer(store)
	if err != nil{
		log.Fatal("Handler not Found")
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	   }
	log.Fatal(srv.ListenAndServe())
}

	func newDBFromConfig() string{
		dbParams := " " + "user=postgres"
		dbParams += " " + "host=localhost"
		dbParams += " " + "port=5432"
		dbParams += " " + "dbname=practice"
		dbParams += " " + "password=password"
		dbParams += " " + "sslmode=disable"

		return dbParams
	}
