package main

import (
	"bufio"
	"chat/utils"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	var username string
	fmt.Scan(&username)
	login(username, conn)
	
	go func() {
		for {
			buf, err := utils.ReadAll(conn)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(buf))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if err := utils.WriteAll([]byte(input), conn); err != nil {
			panic(err)
		}
	}
}

func login(username string, conn net.Conn) {
	if err := utils.WriteAll([]byte(username), conn); err != nil {
		panic(err)
	}

	respBuf, err := utils.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	resp := string(respBuf)
	fmt.Println(resp)
}
