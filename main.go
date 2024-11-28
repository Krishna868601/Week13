package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Database initialization function
func initDB() (*sql.DB, error) {
	// Database connection string (adjust with your actual MySQL credentials)
	dsn := "root:JvNani@5148@tcp(127.0.0.1:3306)/toronto_time" // Replace with your database connection string
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err) // lowercase error message
	}

	// Ping the database to check if the connection works
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err) // lowercase error message
	}

	fmt.Println("Successfully connected to the database")
	return db, nil
}

// Handler to get the current Toronto time
func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	// Get the current time in Toronto timezone
	torontoTime := time.Now().In(time.FixedZone("EST", -5*60*60))

	// Log the current time to the database
	err := logTimeToDB(torontoTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to log time: %v", err), http.StatusInternalServerError) // lowercase error message
		return
	}

	// Return the time as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"current_time": "%s"}`, torontoTime.Format(time.RFC3339))
}

// Log the current time to the database
func logTimeToDB(t time.Time) error {
	// Insert the time into the time_log table
	query := "INSERT INTO time_log (timestamp) VALUES (?)"
	_, err := db.Exec(query, t)
	if err != nil {
		return fmt.Errorf("error inserting time into database: %v", err) // lowercase error message
	}
	fmt.Printf("Logged time: %s\n", t)
	return nil
}

// Main function to start the web server and setup handlers
func main() {
	// Initialize the database connection
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err) // lowercase error message
	}

	// Set up the HTTP server and handlers
	http.HandleFunc("/current-time", getCurrentTime)

	// Start the server on port 8090
	fmt.Println("Starting server on port 8090...")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("failed to start server: %v\n", err) // lowercase error message
	}
}
