package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "db-clusterip"
	port     = 3306
	user     = "root"
	password = "admin"
	dbname   = "db"
)

func main() {
	// Connect to the MariaDB database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("DB connection arguments are valid")

	// Open a test connection to the DB
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connection succeeded")

	// Query the DB
	rows, err := db.Query("SELECT hello FROM messages")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Define a slice to hold JSON objects
	var jsonArray []map[string]string

	// Iterate through result rows
	for rows.Next() {
		var hello string
		if err := rows.Scan(&hello); err != nil {
			log.Fatal(err)
		}

		jsonObject := make(map[string]string)
		jsonObject["hello"] = hello

		jsonArray = append(jsonArray, jsonObject)
	}

	// Get any error encountered during iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Convert the JSON array to a JSON string
	jsonData, err := json.Marshal(jsonArray)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON data
	fmt.Println(string(jsonData))


// Run the Web server and server the results in JSON

	// Define a handler function for the GET request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the CORS headers to allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Encode the response as JSON
		jsonResponse, err := json.Marshal(jsonArray)
		// Check for the errors
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the JSON response
		w.Write(jsonResponse)
	})

	// Start the server on port 8080
	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}