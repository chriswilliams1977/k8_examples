package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//Handles HTTP request
	//http.HandleFunc, tells the http package to handle all requests to the web root ("/") with handler.
	http.HandleFunc("/", handler)

	//Set default port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on localhost:%s", port)
	//http.ListenAndServe, specifying that it should listen on port 8080 on any interface
	//ListenAndServe always returns an error, since it only returns when an unexpected error occurs.
	////In order to log that error we wrap the function call with log.Fatal.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

//Handle request
//The function handler is of the type http.HandlerFunc. It takes an http.ResponseWriter and an http.Request as its arguments.
//An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.
//An http.Request is a data structure that represents the client HTTP request. r.URL.Path is the path component of the request URL.
////The trailing [1:] means "create a sub-slice of Path from the 1st character to the end." This drops the leading "/" from the path name.
func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "World V2"
	}
	fmt.Fprintf(w, "Hello %s!\n", message)
}