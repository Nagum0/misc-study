package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"tcp/utils"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Server listening on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	connReader := bufio.NewReader(conn)
	var length uint32
	binary.Read(connReader, binary.BigEndian, &length)
	yamlBytes := make([]byte, length)
	io.ReadFull(connReader, yamlBytes)
	
	var calcData utils.CalcData
	calcData.FromYaml(yamlBytes)
	fmt.Printf("%v\n", string(yamlBytes))
	fmt.Printf("Received operation: %v\n", calcData)

	result := calcData.Calc()
	var response []byte
	response = fmt.Appendf(response, "Result: %v\n", result)
	connWriter := bufio.NewWriter(conn)
	connWriter.Write(response)
	connWriter.Flush()
}
