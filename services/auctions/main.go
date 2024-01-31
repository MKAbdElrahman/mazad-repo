package main

import (
	"database/sql"
	"fmt"
	"log"
	"mazad/services/auctions/handler"
	"mazad/services/auctions/store"
	"net/http"

	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection parameters
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "user"
	dbPassword := "password"
	dbName := "auctions"

	// Create a PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	auctionStore := store.NewPostgresAuctionStore(db)
	auctionHandler := handler.NewAuctionHandler(auctionStore)

	r.Post("/api/auctions", auctionHandler.HandleCreateAuction)

	http.ListenAndServe(":3000", r)
}
