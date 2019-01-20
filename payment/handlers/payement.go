package handlers

import (
	"auth-payment-microservice/payment/managers"
	"auth-payment-microservice/payment/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index : handlerFunction for root URL
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome in the payment gateway!")
}

// HandlePayments : handlerFunction for /payments/ url path
func HandlePayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	payments := managers.GetPayments()

	json.NewEncoder(w).Encode(payments)
}

// HandlePayment : handlerFunction for /payements/{id} url path
func HandlePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	paymentID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	switch r.Method {

	case "GET":

		payementItem := managers.GetPaymentByID(paymentID)
		json.NewEncoder(w).Encode(payementItem)

	case "DELETE":
		json.NewEncoder(w).Encode(models.Result{
			Message: fmt.Sprintf("HTTP %s Method", r.Method),
		})

	case "POST":
		json.NewEncoder(w).Encode(models.Result{
			Message: fmt.Sprintf("HTTP %s Method ", r.Method),
		})
	}
}
