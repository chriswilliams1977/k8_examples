package main

import (
	"fmt"
	"github.com/chriswilliams1977/gomysqlexample/models"
	"net/http"
	"os"
)

func init(){

}

func init() {

	mySQLUser := os.Getenv("MYSQL_USER")
	mySQLPWD := os.Getenv("MYSQL_PASSWORD")
	mySQLHost := os.Getenv("MYSQL_HOST")
	mySQLDb := os.Getenv("MYSQL_DATABASE")

	models.InitDB(mySQLUser+":"+mySQLPWD+"@tcp("+mySQLHost+":3306)/"+mySQLDb, mySQLDb)
}

func main() {

	http.HandleFunc("/", handler)

	http.HandleFunc("/employees", employeeIndex)
	http.ListenAndServe(":8080", nil)

}

func employeeIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("error")
		http.Error(w, http.StatusText(405), 405)
		return
	}

	employees, err := models.AllEmployees()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, em := range employees {
		fmt.Fprintf(w, "%v, %s, %s\n", em.Id, em.Firstname, em.Lastname)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	target := os.Getenv("TARGET")
	if r.URL.Path[1:] != "" {
		target = r.URL.Path[1:]
	}
	if target == "" {
		target = "World V2"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}

