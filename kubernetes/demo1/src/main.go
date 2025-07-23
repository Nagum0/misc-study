package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Message struct {
	Misc   MiscData   `yaml:"misc"`
	School SchoolData `yaml:"school"`
}

type MiscData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

type SchoolData struct {
	Type        string `yaml:"type"`
	Institution string `yaml:"institution"`
}

func main() {
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/post", postHandler)

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(res, "Well hello there user!\n")
}

func postHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	reqBodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	fmt.Println(string(reqBodyBytes))

	var message Message
	if err := yaml.Unmarshal(reqBodyBytes, &message); err != nil {
		http.Error(res, "Corrupted yaml format", http.StatusBadRequest)
		return
	}

	fmt.Println("RECEIVED DATA: ", message.Misc.Name, message.Misc.Age, message.School.Type)
}
