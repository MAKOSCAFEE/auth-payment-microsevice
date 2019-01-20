package main

import (
	"log"
	"net/http"

	"./routers"
)

func main() {
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
