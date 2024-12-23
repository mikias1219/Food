package main

import (
	"backend/image"
	"backend/login"
	"backend/notify"
	"backend/payment"
	"backend/signup"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request
	var paymentRequest payment.ChapaPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Chapa API Key
	apiKey := "CHASECK_TEST-GdLwnDgIq11gkolK98FSWSjgJSMSamrY" // Replace with your actual API key

	// Initialize payment
	checkoutURL, err := payment.InitiatePayment(apiKey, paymentRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Payment initialization failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the checkout URL
	response := map[string]string{"checkout_url": checkoutURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func connectToDatabase() (*sql.DB, error) {
	dbHost := "postgres"
	dbPort := "5432"
	dbUser := "miki"
	dbPassword := "1219"
	dbName := "miki"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	var db *sql.DB
	var err error

	// Retry logic for database connection
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			if err = db.Ping(); err == nil {
				return db, nil
			}
		}
		log.Printf("Retrying database connection in 5 seconds (%d/10)...", i+1)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to database after retries: %w", err)
}

func main() {
	// Connect to the database
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Serve static files from the uploads directory
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// Initialize routes with database connection
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		notify.HandleEvent(db, w, r)
	})
	http.HandleFunc("/login", login.LoginHandler(db))
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		image.FileUploadHandler(db, w, r)
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		signup.RegisterHandler(w, r, db)
	})
	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		signup.VerifyHandler(w, r, db)
	})
	http.HandleFunc("/payment", PaymentHandler)
	// Start the server
	log.Println("Starting server on http://localhost:8085")
	err = http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal(err)
	}
}
