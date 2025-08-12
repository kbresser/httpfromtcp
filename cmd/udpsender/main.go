package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr := "localhost:42069"
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving UDP address: %s\n", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error dialing UDP: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Sending to %s. Type your message and press Enter to send. Press Ctrl+C to exit.\n", serverAddr)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error sending input: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Message sent: %s", str)
	}
}
