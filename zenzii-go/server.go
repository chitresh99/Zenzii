// server.go
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func server() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port 9000...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	// Create a new file to save received content
	outFile, err := os.Create("received_file")
	if err != nil {
		fmt.Println("File creation error:", err)
		return
	}
	defer outFile.Close()

	// Copy data from connection to file
	bytesWritten, err := io.Copy(outFile, conn)
	if err != nil {
		fmt.Println("Error receiving file:", err)
		return
	}
	fmt.Printf("File received successfully (%d bytes)\n", bytesWritten)
}
