// client.go
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func client() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	file, err := os.Open("file_to_send")
	if err != nil {
		fmt.Println("File open error:", err)
		return
	}
	defer file.Close()

	// Send file over the connection
	bytesSent, err := io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
		return
	}
	fmt.Printf("File sent successfully (%d bytes)\n", bytesSent)
}
