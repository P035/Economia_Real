package main

import (

	"fmt"
	"net"
	"os"
	"bufio"
	"github.com/P035/Economia_Real/db"
)

// This two constants will be the data that I'll pass to the TCPAddr struct.
const ip = "192.168.1.8"
const port = 3000

func main() {

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

	/*
	
	Test1, this test will check if the tcp connection works, by accepting only one connection and just printing on the screen what information is being recived from the client.
	The client that i'm going to use is telnet.
	
	<CODE>

	conn, err := listener.Accept()
	// Check if it had an error accepting the connection.
	if err != nil {

		fmt.Println("Error accepting connection:", err)
		os.Exit(1)
	}
	fmt.Println("Connection accepted from", conn.RemoteAddr().String())
	defer conn.Close()

	// It crates a reader, so it can use it for reading from the connection
	reader := bufio.NewReader(conn)

	// Loop of reading from the client until the connection closes

	for {

		// Read from the connection with ioutil.ReadAll.
		data, err := reader.ReadBytes('\n')
		// Check for errors
		if err != nil {

			fmt.Println("Error:", err)
			break
		}
		fmt.Println(string(data[:len(data) - 1]))
	}
	<CODE>
	*/
	// Accept connection and check for errors
	conn, err := listener.Accept()
	if err != nil {

		fmt.Println("Error accepting connecton:", err)
		os.Exit(1)
	}
	fmt.Println("Connection created with", conn.RemoteAddr())
	defer conn.Close()

	// It will read the user and the password from the client so it needs to create a reader
	reader := bufio.NewReader(conn)

	// It has to send two messages to the client, so here it creates the two string messages, but in bytes.
	user_msg := []byte("Enter your username: ")
	psw_msg := []byte("Enter your password: ")
	conn.Write(user_msg)
	user, err := reader.ReadBytes('\n')
	if err != nil {

		fmt.Println("Error reading user name:", err)
		os.Exit(1)
	}
	conn.Write(psw_msg)
	psw, err := reader.ReadBytes('\n')
	if err != nil {

		fmt.Println("Error reading password:", err)
	}
	db.Init()
	db.Select([]byte("SELECT * FROM usuarios WHERE Nombre = 'peo' AND Contraseña = 'Culo'"))
	fmt.Println(string(user) + string(psw))
}
