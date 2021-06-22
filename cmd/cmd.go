// This packages will contain all the commands that the economy game is going to have.
package cmd

import (

	"net"
	"github.com/P035/Economia_Real/db"
)

func login(conn net.Conn) {

	usr_msg := []byte("Enter your username: ")
	_, err := conn.Write(usr_msg)
	if err != nil {

		fmt.Println("Error sending message to client")
		return
	}
	psw_msg := []byte("Enter your password: ")
	_, err := conn.Write(psw_msg)
}
