// This packages will contain all the commands that the economy game is going to have.
package cmd

import (

	"net"
	"fmt"
	"bufio"
	"github.com/P035/Economia_Real/db"
)

// This function will read from the client the username and the password and then search for those credentials inthe db
func Login(conn net.Conn) []db.Usuario{

	// Initialize the database
	db.Init()

	// It's going to use a bufio reader for read from the client
	rdr := bufio.NewReader(conn)

	// After sending the messages it will read from the client
	usr_msg := []byte("Enter your username: ")
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
	fmt.Println("User:", string(usr))

	psw_msg := []byte("Enter your password: ")
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
	fmt.Println("Password:", string(psw))

	// If it successfully read the data from the client it is going to see if there is an user with that password in the database
	query := "SELECT * FROM usuarios WHERE Nombre = '" + string(usr[:len(usr) - 2]) + "' AND Contraseña = '" + string(psw[:len(psw) - 2]) + "';"
	data := db.Select(query)
	return data
}
