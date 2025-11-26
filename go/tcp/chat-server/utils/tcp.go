package utils

import (
	"encoding/binary"
	"io"
	"net"
)

func ReadAll(conn net.Conn) ([]byte, error) {
	var size uint32
	if err := binary.Read(conn, binary.BigEndian, &size); err != nil {
		return nil, err
	}
	
	buf := make([]byte, size)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func WriteAll(data []byte, conn net.Conn) error {
	if err := binary.Write(conn, binary.BigEndian, uint32(len(data))); err != nil {
		return err
	}
	
	if _, err := conn.Write(data); err != nil {
		return err
	}

	return nil
}
