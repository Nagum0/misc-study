package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHandler)
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "-- RESPONSE FROM LOADBALANCER SERVER")
}
