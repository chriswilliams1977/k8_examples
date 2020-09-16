package main

import(
	"fmt"
	"github.com/chriswilliams1977/lifecycle/datafile"
	"log"
	"net/http"
	"os"
)

func main(){
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

func handler(w http.ResponseWriter, r *http.Request) {
	//Get data from file at /tmp/names.txt
	records, counts := datafile.GetDataUsingSlice("/tmp/names.txt")
	//print names and count
	for i, record := range records{
		//prints to logs
		fmt.Printf("%s: %d\n", record, counts[i])
		//prints to window
		fmt.Fprintf(w, "%s: %d\n", record, counts[i])
	}
}