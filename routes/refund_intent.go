package routes

import (
	"net/http"

	"github.com/tej112/helpers"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/refund"
)

func RefundIntent(w http.ResponseWriter, r *http.Request) {

	// Get the id from the URL
	id := mux.Vars(r)["id"]
	if id == "" {
		helpers.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	// Create a new refund for the payment intent
	_, err := refund.New(&stripe.RefundParams{
		PaymentIntent: stripe.String(id),
	})

	if err != nil {
		e := err.(*stripe.Error)
		helpers.WriteJSON(w, e.HTTPStatusCode, map[string]string{"error": "Error refunding payment intent", "message": e.Msg})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, map[string]string{"message": "Intent refunded successfully"})
}
