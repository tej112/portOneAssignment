package routes

import (
	"net/http"

	"github.com/tej112/helpers"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

func ListIntents(w http.ResponseWriter, r *http.Request) {
	// List all the payment intents
	params := &stripe.PaymentIntentListParams{}

	iter := paymentintent.List(params)

	// Check if there is an error
	if iter.Err() != nil {
		helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error listing payment intents"})
		return
	}

	// Iterate over the payment intents and append them to the slice
	itents := []stripe.PaymentIntent{}
	for iter.Next() {
		itents = append(itents, *iter.PaymentIntent())
	}

	// Write the response
	helpers.WriteJSON(w, http.StatusOK, itents)

}
