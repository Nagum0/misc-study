package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error happened: ", err.Error())
	}
	defer f.Close()

	_, err = f.Write([]byte("I like Makima!\n"))
	if err != nil {
		fmt.Println("Error happened: ", err.Error())
	}
}
