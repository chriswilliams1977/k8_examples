package main

import(
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
		port = "8081"
	}

	log.Printf("Listening on localhost:%s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a list of banks")
}