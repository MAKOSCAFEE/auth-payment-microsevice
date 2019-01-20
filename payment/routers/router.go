package routers

import (
	"net/http"

	"auth-payment-microservice/payment/handlers"

	"github.com/gorilla/mux"
)

// Route : Route structure
type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes []Route

// init method which will be automatically run
func init() {

	// init root URL and handler function
	routes = append(routes, Route{
		Name:        "Index",
		Methods:     []string{"GET"},
		Pattern:     "/",
		HandlerFunc: handlers.Index,
	})

	/* init /payments/{id} URL and handler function.
	One handler function for three type of requests */
	routes = append(routes, Route{
		Name:        "Payment",
		Methods:     []string{"DELETE", "GET", "POST"},
		Pattern:     "/payments/{id}",
		HandlerFunc: handlers.HandlePayment,
	})

	/* init /payments URL and handler function. */
	routes = append(routes, Route{
		Name:        "Payments",
		Methods:     []string{"GET"},
		Pattern:     "/payments",
		HandlerFunc: handlers.HandlePayments,
	})
}

// NewRouter : registering all url paths
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
