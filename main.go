package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// establish the server
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err.Error())
		os.Exit(1)
	}
	// create the connection and listen to the port
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to the port", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Ready to recieve messages on port 8888")

	for {
		// create a buffer and read from the messages to the port
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from connection", err.Error())
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", addr, message)
		// upcase the message and return it from the sending port
		message = strings.ToUpper(message)
		_, err = conn.WriteToUDP([]byte(message), addr)
		if err != nil {
			fmt.Println("Error sending the message:", err.Error())

		}
	}

}
