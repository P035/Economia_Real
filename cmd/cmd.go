// This packages will contain all the commands that the economy game is going to have.
package cmd

import (

	"net"
	"fmt"
	"bufio"
	"github.com/P035/Economia_Real/db"
)

var usr_msg []byte = []byte("Enter your username: ")
var psw_msg []byte = []byte("Enter your password: ")

// This function will read from the client the username and the password. And after that it's going to register the data into the database
// The return value will mean if the function succeeded.
func Register(conn net.Conn) []db.Usuario {

	// Initialize the database
	db.Init()

	// For reading from the client is going to use a bufio reader
	rdr := bufio.NewReader(conn)

	// Every time before reading the server will send a message like 'Enter your username: '

	// First is going to read the username.
	_, err := conn.Write(usr_msg)
	// Check for errors
	if err != nil {

		fmt.Println("Error sending message to the client:", err)
		return nil
	}
	usr, err := rdr.ReadBytes('\n')
	if err != nil {

		fmt.Println("Error reading from the client:", err)
		return nil
	}
	fmt.Println("Username:", string(usr[:len(usr) - 1]))

	// Now is going to read the password.
	_, err = conn.Write(psw_msg)
	// Check for errors
	if err != nil {

		fmt.Println("Error sending message to the client:", err)
		return nil
	}
	psw, err := rdr.ReadBytes('\n')
	if err != nil {

		fmt.Println("Error reading from the client:", err)
		return nil
	}
	fmt.Println("Password:", string(psw[:len(psw) - 1]))

	// After reading all the data from the client it has to check if it exists in the database.
	query := "SELECT * FROM usuarios WHERE Nombre = '" + string(usr[:len(usr) - 2]) + "';"
	data := db.Select(query)
	if len(data) > 0 {

		conn.Write([]byte("Error: The username allready exists!\n"))
		return nil
	}

	// The query is going to be an insert query.
	query = "INSERT INTO usuarios(Nombre, Contraseña) VALUES('" + string(usr[:len(usr) - 2]) + "', '" + string(psw[:len(psw) - 2])  + "');"
	result := db.Insert(query)
	fmt.Println("Result:", result)
	return result
}

// This function will read from the client the username and the password and then search for those credentials inthe db
func Login(conn net.Conn) []db.Usuario{

	// Initialize the database
	db.Init()

	// It's going to use a bufio reader for read from the client
	rdr := bufio.NewReader(conn)

	// After sending the messages it will read from the client
	_, err := conn.Write(usr_msg)
	if err != nil {

		fmt.Println("Error sending message to client:", err)
		return nil
	}

	// If there is no errors it will read the username
	usr, err := rdr.ReadBytes('\n')
	if err != nil {

		fmt.Println("Error reading from client:", err)
		return nil
	}
	fmt.Println("User:", string(usr[:len(usr) - 1]))

	_, err = conn.Write(psw_msg)
	if err != nil {

		fmt.Println("Error sending message to client:", err)
		return nil
	}
	// If there is no errors it will read the password
	psw, err := rdr.ReadBytes('\n')
	if err != nil {

		fmt.Println("Error reading from client:", err)
		return nil
	}
	fmt.Println("Password:", string(psw[:len(psw) - 1]))

	// If it successfully read the data from the client it is going to see if there is an user with that password in the database
	query := "SELECT * FROM usuarios WHERE Nombre = '" + string(usr[:len(usr) - 2]) + "' AND Contraseña = '" + string(psw[:len(psw) - 2]) + "';"
	data := db.Select(query)
	return data
}
