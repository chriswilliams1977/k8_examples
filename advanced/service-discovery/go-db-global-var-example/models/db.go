//The `defer` statement is used to run some code just immediately before a function returns.
//So calling `defer db.Close()` will close the database connection pool when your `InitDB` function returns.
//The right way to handle this is have `InitDB` return the `*sql.DB` object and then defer the close from within your main.main() function.
//Or alternatively, you can just remove it like you have done.
//If the database connection pool is to exist the whole time your app is running then it's fine to not close it,
//it will be removed from memory when your Go app is terminated. Generally, it's only really necessary to call db.Close()
//in things where the connection pool is short-lived, such as a test.
package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//make sure you db connection is global so other func calling db can reference the connection
var db *sql.DB
var err error

func InitDB(dataSourceName string, dbName string) {

	// Configure the database connection
	db, err = sql.Open("mysql", dataSourceName)
	ErrorCheck(err)

	// Initialize the first connection to the database, to see if everything works correctly.
	//PingDB(db)

	//Create Database using Exec. Exec is used for queries where no rows are returned.
	_,err = db.Exec( "CREATE DATABASE IF NOT EXISTS "+dbName)
	ErrorCheck(err)
	fmt.Println("Successfully created database..")

	//Choose the correct database each time you start a MySQL session with USE.
	_,err = db.Exec("USE "+dbName+"")
	ErrorCheck(err)
	fmt.Println("DB selected successfully..")

	//Prepare creates a prepared statement for later queries or executions.
	stmt, err := db.Prepare("CREATE Table IF NOT EXISTS "+dbName+"(id int NOT NULL AUTO_INCREMENT, firstname varchar(50), lastname varchar(30), PRIMARY KEY (id));")
	ErrorCheck(err)

	_, err = stmt.Exec()
	ErrorCheck(err)
	fmt.Println("Table created successfully..")

	populateDB(db, dbName)

	//close the Database.
	//defer db.Close()
}

func populateDB(db *sql.DB, dbName string){
	fmt.Println("Populating database")

	//insert into db.
	stmt, e := db.Prepare("INSERT INTO "+dbName+" (firstname, lastname) VALUES ('Chris', 'Williams')")
	ErrorCheck(e)

	//get result, last entry added
	res, e := stmt.Exec()
	ErrorCheck(e)

	//get result, last entry added
	id, e := res.LastInsertId()
	ErrorCheck(e)

	fmt.Println("Insert id", id)

}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}