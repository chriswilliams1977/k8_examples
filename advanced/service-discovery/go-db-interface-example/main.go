package main

import (
"fmt"
"github.com/chriswilliams1977/gomysqlinterfaceexample/models"
"log"
"net/http"
"os"
)

var mySQLUser = os.Getenv("MYSQL_USER")
var mySQLPWD = os.Getenv("MYSQL_PASSWORD")
var mySQLHost = os.Getenv("MYSQL_HOST")
var mySQLDb = os.Getenv("MYSQL_DATABASE")

type Env struct {
	db models.Datastore
}

func main(){
db, err := models.NewDB(mySQLUser+":"+mySQLPWD+"@tcp("+mySQLHost+":3306)/"+mySQLDb, mySQLDb)
if err != nil {
log.Panic(err)
}
env := &Env{db: db}

http.HandleFunc("/employees", env.employeeIndex)
http.ListenAndServe(":8080", nil)

}

func (env *Env) employeeIndex(w http.ResponseWriter, r *http.Request) {
if r.Method != "GET" {
fmt.Println("error")
http.Error(w, http.StatusText(405), 405)
return
}

employees, err := env.db.AllEmployees()
if err != nil {
http.Error(w, http.StatusText(500), 500)
return
}
for _, em := range employees {
fmt.Fprintf(w, "%v, %s, %s\n", em.Id, em.Firstname, em.Lastname)
}
}