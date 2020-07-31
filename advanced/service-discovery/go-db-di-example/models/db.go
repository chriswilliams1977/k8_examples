package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(dataSourceName string, dbName string) (*sql.DB, error) {

	// Configure the database connection
	db, err := sql.Open("mysql", dataSourceName)
	ErrorCheck(err)

	// Initialize the first connection to the database, to see if everything works correctly.
	PingDB(db)

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
	return db, nil
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