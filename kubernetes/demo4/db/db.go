package main

import (
	"fmt"
	"log"
	"net/http"
)

func dataHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello, World from db-service!\n")
}

func main() {
	http.HandleFunc("/data", dataHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
