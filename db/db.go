package db

import (

	"os"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type usuario struct {

	id int
	name string
	psw string
}

var db *sql.DB = nil

// This function will connect to the database, so it is like an initializer.
func Init() {
	// First it has to connect to my "server", that is actually just localhost.
	conn, err := sql.Open("mysql", "root:Naranjo7854@/relacional")
	// Check for errors
	if err != nil {

		fmt.Println("Error establishing connection with database", err)
		os.Exit(1)
	}
	db = conn
	fmt.Println("Connection created with database")
}

// I'm going to try doing a select query without a transaction
func Select(query []byte) []usuario{

	for _, i := range query {

		fmt.Println(string(i))
	}
	data, err := db.Query(string(query))
	if err != nil{

		fmt.Println("Error reading from database:", err)
		os.Exit(1)
	}
	var aux bool
	var arr[]usuario
	var user usuario
	for {

		aux = data.Next()
		if aux == false {

			fmt.Println(arr)
			return arr
		}else {

			data.Scan(&user.id, &user.name, &user.psw)
			arr = append(arr, user)
			fmt.Println(arr)
		}
	}
}

// The insert query with the database/sql library works if you do it
// directly with the database (without a transaction).
func Insert(query string) {

	result, err := db.Exec(query)
	if err != nil {

		fmt.Println("Error executing query", err)
		os.Exit(1)
	}
	fmt.Println("Query executed successfully")
	id, _ := result.LastInsertId()
	fmt.Println("Last insert id:", id)
}
