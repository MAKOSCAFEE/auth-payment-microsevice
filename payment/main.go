package main

import (
	"log"
	"net/http"

	"auth-payment-microservice/payment/routers"
)

func main() {
	addr := ":5000"
	router := routers.NewRouter()
	log.Println("Server has started at http://localhost:", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
