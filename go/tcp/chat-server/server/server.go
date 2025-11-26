package main

import (
	"chat/utils"
	"fmt"
	"net"
	"sync"
)

type User struct {
	Name string
	RemoteAddr net.Addr
	Conn net.Conn
}

var users map[string]User
var usersMutex sync.RWMutex

func main() {
	users = map[string]User{}

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	user := login(conn)

	for {
		req, err := utils.ReadAll(conn)
		if err != nil {
			fmt.Printf("Client %v diconnected.\n", user.Name)
			usersMutex.Lock()
			delete(users, user.Name)
			usersMutex.Unlock()
			return
		}
		fmt.Printf("Message from %v: %v\n", user.Name, string(req))
		
		usersMutex.RLock()
		for username, otherUser := range users {
			if user.Name == otherUser.Name {
				continue
			}

			if err := utils.WriteAll(req, otherUser.Conn); err != nil {
				fmt.Printf("Error while writing to %v\n", username)
				continue
			}
		}
		usersMutex.RUnlock()
	}
}

func login(conn net.Conn) User {
	fmt.Printf("New connection from: %v\n", conn.RemoteAddr())

	usernameBuf, err := utils.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	user := User{
		Name: string(usernameBuf),
		RemoteAddr: conn.RemoteAddr(),
		Conn: conn,
	}
	usersMutex.Lock()
	users[user.Name] = user
	usersMutex.Unlock()

	resp := fmt.Sprintf("Logged in as: %v", user.Name)
	if err := utils.WriteAll([]byte(resp), conn); err != nil {
		panic(err)
	}
	fmt.Printf("%v logged in as %v\n", user.RemoteAddr, user.Name)

	return user
}
