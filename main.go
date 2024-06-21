package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tej112/middleware"
	"github.com/tej112/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v78"
)

var PORT string

func init() {
	// Load the .env file
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	// Get The STRIPE_SECRET_KEY from the .env file
	StripeKey := os.Getenv("STRIPE_SECRET_KEY")
	if StripeKey == "" {
		log.Fatal("Error loading STRIPE_SECRET_KEY from .env file")
	}

	// Set the STRIPE_SECRET_KEY
	stripe.Key = StripeKey

	// Get the PORT from the .env file
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Use the middleware for logging and setting up the headers
	router.Use(middleware.Logger, middleware.HeaderSetup)

	// Create group with prefix /api
	api := router.PathPrefix("/api").Subrouter()

	// Create subrouter for version 1
	v1 := api.PathPrefix("/v1").Subrouter()

	// Create intent for payment
	v1.HandleFunc("/create_intent", routes.CreateIntent).Methods("POST")

	// Capture the created intent
	v1.HandleFunc("/capture_intent/{id}", routes.CaptureIntent).Methods("POST")

	// Create a refund for the created intent
	v1.HandleFunc("/create_refund/{id}", routes.RefundIntent).Methods("POST")

	// Get a List of all intents
	v1.HandleFunc("/get_intents", routes.ListIntents).Methods("GET")

	// Start the server
	log.Println("Server started on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
