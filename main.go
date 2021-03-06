package main

import (
	handler "github/golang_practice/handler"
	"github/golang_practice/storage/postgres"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
)
func main (){
    

	session := sessions.NewCookieStore([]byte("my_secret"))


	newDbString := newDBFromConfig()

	store, err := postgres.NewStorage(newDbString)
	if err != nil {

		log.Println("error db")
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	
	db, err := sqlx.Connect("postgres", newDbString)
	if err != nil {
		return 
	}

	   
	r, err := handler.NewServer(store, decoder, session, db)
	if err != nil {
		log.Println("error on handelr")


	   decoder := schema.NewDecoder()
	   decoder.IgnoreUnknownKeys(true)




	   session := sessions.NewCookieStore([]byte("1234"))

	   
	r ,err := handler.NewServer(store, decoder, session)
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
