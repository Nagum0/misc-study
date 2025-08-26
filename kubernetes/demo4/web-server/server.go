package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func dataHandler(resp http.ResponseWriter, req *http.Request) {
	r, err := http.Get("http://db-service:8081/data")
	if err != nil {
		fmt.Fprintf(resp, "Error while getting the data...\n")
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(resp, "Error while reading the data...\n")
	}

	fmt.Fprintf(resp, "Data:\n%v", string(body))
}

func main() {
	http.HandleFunc("/data", dataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
