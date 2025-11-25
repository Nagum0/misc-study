package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"tcp/utils"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	calcData := utils.CalcData{
		A: 42,
		Op: utils.PLUS,
		B: 25,
	}
	calcDataYaml := calcData.ToYaml()
	calcDataYamlLength := len(calcDataYaml)
	connWriter := bufio.NewWriter(conn)
	binary.Write(connWriter, binary.BigEndian, uint32(calcDataYamlLength))
	connWriter.Write(calcDataYaml)
	connWriter.Flush()
	fmt.Printf("Sent operation: %v\n", calcData)

	connReader := bufio.NewReader(conn)
	response, _ := connReader.ReadString('\n')
	fmt.Println(response)
}
