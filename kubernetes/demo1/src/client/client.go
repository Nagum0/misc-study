package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Expected 1 arguemtn\n")
		os.Exit(1)
	}

	yamlContent, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	body := bytes.NewReader(yamlContent)

	res, err := http.Post("http://192.168.49.2:30000/post", "application/yaml", body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(res.Body)
}
