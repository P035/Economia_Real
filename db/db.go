package db

import (

	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type Usuario struct {

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
	}
	db = conn
	fmt.Println("Connection created with database")
}

// I'm going to try doing a select query without a transaction
func Select(query string) []Usuario{

	data, err := db.Query(query)
	if err != nil{

		fmt.Println("Error reading from database:", err)
		return nil
	}
	var aux bool
	var arr[]Usuario
	var user Usuario
	for {

		aux = data.Next()
		if aux == false {

			return arr
		}else {

			data.Scan(&user.id, &user.name, &user.psw)
			arr = append(arr, user)
		}
	}
}

// The insert query with the database/sql library works if you do it
// directly with the database (without a transaction).
// The return value will deppend if the query succseed or not.
func Insert(query string) bool{

	result, err := db.Exec(query)
	if err != nil {

		fmt.Println("Error executing query", err)
		return false
	}
	fmt.Println("Query executed successfully")
	id, _ := result.LastInsertId()
	fmt.Println("Last insert id:", id)
	return true
}
