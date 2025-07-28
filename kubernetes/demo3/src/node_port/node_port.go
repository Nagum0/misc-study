package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHandler)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "-- RESPONSE FROM NODEPORT SERVER")
}
