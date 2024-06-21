package routes

import (
	"encoding/json"
	"net/http"

	"github.com/tej112/helpers"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

// Capture Intent request body
type CaptureIntentRequest struct {
	AmountToCapture float64 `json:"amount_to_capture"`
}

func CaptureIntent(w http.ResponseWriter, r *http.Request) {
	// Get the id from the URL
	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Load the request body
	var req CaptureIntentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate the request
	if req.AmountToCapture <= 0 {
		http.Error(w, "Invalid amount to capture", http.StatusBadRequest)
		return
	}

	// CaptureInternt Params
	params := &stripe.PaymentIntentCaptureParams{
		AmountToCapture: stripe.Int64(int64(req.AmountToCapture * 100)), // Amount is converted to whole number
	}

	// Capture the intent
	_, err = paymentintent.Capture(id, params)
	if err != nil {
		// Convert the error to stripe error to get the HTTP status code and message
		e := err.(*stripe.Error)

		// Write the error response
		helpers.WriteJSON(w, e.HTTPStatusCode, map[string]string{"error": "Error capturing payment intent", "message": e.Msg})
		return
	}

	// Write the success response
	helpers.WriteJSON(w, http.StatusOK, map[string]string{"message": "Intent captured successfully"})
}
