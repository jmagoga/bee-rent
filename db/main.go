// package db

// import (
// 	"log"
// 	"os"

// 	"github.com/go-pg/pg/v10"
// )

// func Connect() {
// 	opts := &pg.Options{
// 		User:     "user",
// 		Password: "secret",
// 		Addr:     "localhost:5432",
// 		Database: "meetmeup",
// 		// e o nome da DB?
// 	}
// 	var db *pg.DB = pg.Connect(opts)
// 	if db == nil {
// 		log.Printf("Failed to connect to DB")
// 		os.Exit(100)
// 	}
// 	log.Printf("Connection to DB succesful!")

// 	// once done with DB, close it.
// 	closeErr := db.Close()
// 	if closeErr != nil {
// 		log.Printf("Error while closing the connection. Reason: , %v\n", closeErr)
// 		os.Exit(100)
// 	}
// 	log.Printf("Connection closed succesfully.\n")
// 	return
// }
