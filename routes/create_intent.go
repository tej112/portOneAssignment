package routes

import (
	"encoding/json"
	"net/http"

	"github.com/tej112/helpers"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

// Create Intent request body to get the amount and currency
type CreateIntentRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func CreateIntent(w http.ResponseWriter, r *http.Request) {
	// load the request body
	var req CreateIntentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	// validate the request
	if req.Amount <= 0 {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid amount"})
		return
	}
	if !helpers.CheckCurrency(req.Currency) {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid currency"})
		return
	}

	// Payment Intent Params
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(int64(req.Amount * 100)),                          // Amount is converted to whole number
		Currency:      stripe.String(string(req.Currency)),                            // Currency
		CaptureMethod: stripe.String(string(stripe.PaymentIntentCaptureMethodManual)), // Manual capture to capture the payment later
		PaymentMethod: stripe.String(string("pm_card_visa")),                          // Default payment method
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled:        stripe.Bool(true),      // Enable automatic payment methods
			AllowRedirects: stripe.String("never"), // Disable redirects
		},
		Confirm: stripe.Bool(true), // Confirm the payment intent so that the client can pay and capture the payment later
	}

	// Create the payment intent
	intent, err := paymentintent.New(params)
	if err != nil {
		e := err.(*stripe.Error)
		helpers.WriteJSON(w, e.HTTPStatusCode, map[string]string{"error": "Error creating payment intent", "message": e.Msg})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"client_secret": intent.ClientSecret, "id": intent.ID})
}
