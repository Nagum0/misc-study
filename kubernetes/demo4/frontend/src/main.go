package main

import (
	"log"
	"net/http"
)

func main()  {
	fs := http.FileServer(http.Dir("./src/ui/templ/"))
	http.Handle("/ui/", http.StripPrefix("/ui/", fs))

	log.Println("Listening on port :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
