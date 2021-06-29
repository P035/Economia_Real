package main

import (

	"fmt"
	"net"
	"os"
	"bufio"
	"strconv"
	"github.com/P035/Economia_Real/cmd"
	"github.com/P035/Economia_Real/db"
)

func handle(conn net.Conn) {

	var logged_in bool = false
	var usr []db.Usuario

	// It will use a bufio reader to read from the client.
	rdr := bufio.NewReader(conn)

	for {

		data, err := rdr.ReadBytes('\n')
		if err != nil {

			// Check if there is an EOF error
			if err.Error() == "EOF"{

				fmt.Println("Connection closed with:", conn.RemoteAddr())
				conn.Close()
				break
			}else {

				fmt.Println("Error reading from client:", err)
			}
		}else {


			fmt.Println(conn.RemoteAddr(), ":", string(data[:len(data) - 2]))
			if string(data[:len(data) - 2]) == "+login"{

				if logged_in == false {


					fmt.Println(string(data))
					db_data := cmd.Login(conn)
					if len(db_data) == 0{

						conn.Write([]byte("Username or password incorrect.\n"))
					}else {

						conn.Write([]byte("Welcome " + db_data[0].Name + "\n"))
						logged_in = true
						usr = db_data
						fmt.Println(conn.RemoteAddr(), "=", usr)
					}
				}else {

					conn.Write([]byte("You are allready logged in. Logout first!\n"))
				}
			}else if string(data[:len(data) - 2]) == "+register"{

				if logged_in == false{

					db_data := cmd.Register(conn)
					if db_data == nil {

						conn.Write([]byte("Error registering your account :C\n"))
						fmt.Println(db_data)
					}else {

						conn.Write([]byte("You are now succesfully registered and logged in with your new account.\n"))
						logged_in = true
						usr = db_data
						fmt.Println(conn.RemoteAddr(), "=", usr)
					}
				}else {

					conn.Write([]byte("You are allready logged in. Logout first!\n"))
				}
			}else if string(data[:len(data) - 2]) == "+logout"{

				if logged_in == false {

					conn.Write([]byte("You are allready logged out. Loggin first! \n"))
				}else {

					logged_in = false
					usr = nil
					conn.Write([]byte("You are logged out!"))
				}
			}
		}
	}
}

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

	// After creating the listener it has to accept connections and create a gorroutine for each
	for {

		// Wait and accept connections
		conn, err := listener.Accept()
		// Check for errors
		if err != nil {

			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Connection created with:", conn.RemoteAddr())

		// It creates the corroutine to handle the connection
		go handle(conn)
	}
}
