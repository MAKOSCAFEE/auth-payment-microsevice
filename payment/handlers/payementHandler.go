package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../managers"
	"../models"

	"github.com/gorilla/mux"
)

// Index : handlerFunction for root URL
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome in the payment gateway!")
}

// HandlePayments : handlerFunction for /users/ url path
func HandlePayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := managers.GetPayments()

	json.NewEncoder(w).Encode(users)
}

// HandlePayment : handlerFunction for /user/{id} url path
func HandlePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	paymentID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	switch r.Method {

	case "GET":

		userItem := managers.GetPaymentByID(paymentID)
		json.NewEncoder(w).Encode(userItem)

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
