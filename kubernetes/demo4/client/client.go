package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(5 * time.Second)
		r, err := http.Get("http://demo4-web-server-service:8080/data")
		if err != nil {
			log.Println(err.Error())
			continue
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
		}

		fmt.Println(string(body))
		r.Body.Close()
	}
}
