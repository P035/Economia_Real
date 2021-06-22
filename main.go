package main

import (

	"fmt"
	"net"
	"os"
	"bufio"
	"strconv"
	"github.com/P035/Economia_Real/db"
)

func main() {

	// Read from the terminal arguments.
	if len(os.Args) < 3 {

		fmt.Println("Fogot to specify ip and port. (ip first then port)")
		os.Exit(1)
	}
	ip := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])

	// Declare the addr struct and define the IP and port fields.
	var addr net.TCPAddr
	addr.IP = net.ParseIP(ip)
	addr.Port = port

	fmt.Println("Addr building succed", addr.String())

	// After creating the addres it has to create the tcp listener.
	listener, err := net.ListenTCP("tcp", &addr)
	// Check if it had an error creating the listener
	if err != nil {

		// If there is an error i'm going to print it on the screen an exit the program with os.Exit(1)
		fmt.Println("Error creating listener:", err)
		os.Exit(1)
	}
	// If there isn't errors it defers the close function of the listener.
	fmt.Println("Listener created")
	defer listener.Close()
}
