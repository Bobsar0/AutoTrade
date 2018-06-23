package main

import (
	"log"
	"net/http"
	"github.com/chidi150c/autotrade/app"
	"os"
	hatchuid "github.com/nu7hatch/gouuid"
	"github.com/gorilla/handlers"
	"strings"
)

//Program starts here
func main() {

	s := app.NewSession() //Initializes new session
	h := app.NewAppHandler(s) //Passes the session to initialize a new instance of appHandler
	
	//Starting the getTicker goroutine

	go app.GetTicker(s.GetTickerChan)
	go app.GetBalance(s.GetBalanceChan)
	go app.DBService(s.AddOrUpdateDbChan, s.GetDbChan, s.DeleteDbChan, hatchID())
	
	log.Println("Connecting to server on port 8000...")
	log.Fatalln(http.ListenAndServe(":8000", handlers.CombinedLoggingHandler(os.Stderr, h))) //Set listening port (:8080). Handler is h indicating that chi router is used. log.Fatal checks for error and outputs if any.
}
func hatchID()(chan string){
	c := make(chan string)
	go func (){
		for{
			u4, err := hatchuid.NewV4()
			if err != nil {
				log.Printf("FATAL0: error: %v", err)
			}
			u4hatch := strings.Split(u4.String(), "-")
			c <-u4hatch[0]
		}
	}()
	return c
}