//Using http package to build a simple server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to AutoTrade<h1>") //Prints to webpage
}

func main() {
	http.HandleFunc("/", server)                   //Register the handler function (server) for the given pattern ("/")
	fmt.Println("Connecting to server on port 8000...")
	log.Fatalln(http.ListenAndServe(":8000", nil)) //Set listening port (:8080). Handler is nil indicating that DefaultServeMux should be used. log.Fatal checks for error and outputs if any.
}
