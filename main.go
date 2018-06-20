package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi" //Using chi mux/router
)

//Type appHandler contains the chi mux and session and implements the ServeMux method
type appHandler struct{
	mux *chi.Mux
	session *Session
}

//indexHandler delivers the Home page to the user
func (h *appHandler)indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Welcome to AutoTrade<h1>
		<p><a href="/ticker">ticker</a></p>`) //Prints ticker to webpage
}

//getTickerHandler presents the ticker price to the user
func (h *appHandler)getTickerHandler(w http.ResponseWriter, r *http.Request){
	tickerChan := make(chan float64) //tickerChan represents a channel that returns the ticker price
	h.session.getTickerChan <- tickerChan //Send the content(ticker price) in tickerChan to session
	ticker := <- tickerChan //ticker receives the ticker price via tickerChan
	
	responseToUser := fmt.Sprintf("<h1>Ticker: %.8f<h1>", ticker) //Returns response to user (which contains ticker) as a string
	fmt.Fprintf(w, "%s", responseToUser) //Prints response to user on the web page
	return 
}

//A test function to simulate getting ticker price from a trading site
//In the real world, this will be achieved by communicating with the site's API
func functhatreturnticker()float64{
	return 0.002134442
}

//PASSING CHANNEL OVER CHANNEL
//getTicker() takes in a channel(gtc) which provides a channel of float64 (tickerChan)
//It provides the ticker price via a function that interacts with the API
//Then sends the ticker price back to the caller of the function (HANLDER) through the tickerChan
func getTicker(gtc chan chan float64 ){
	var ticker float64
	for{
		select{
		case tickerChan := <- gtc: //if tickerChan receives a chan of float64 (gtc) from the getTicker() func caller
			ticker = functhatreturnticker() //Call the func to GET the ticker from the trading site (Using a test function for now)
			tickerChan <- ticker //Send the ticker back to the caller function(handler) via the tickerChan
		}
	}
}

//Session carries the state parameters of a particular user
type Session struct{
	getTickerChan chan chan float64
}

//NewSession returns a new instance of *Session 
func NewSession () *Session{
	return &Session{
	}
}

//NewAppHandler returns a new instance of *appHandler
func NewAppHandler (s *Session) *appHandler{
	h := &appHandler{
		mux: chi.NewRouter(),
		session: s,
	}
	h.mux.Get("/ticker", h.getTickerHandler)
	h.mux.Get("/", h.indexHandler)
	return h
}

//appHandler implements ServeHTTP method making it a Handler
func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

//Program starts here
func main() {
	//channel to communicate with getTicker goroutine
	getTickerChan := make(chan chan float64)

	s := NewSession() //Initializes new session
	h := NewAppHandler(s) //Passes the session to new instance of appHandler
	
	h.session.getTickerChan = getTickerChan
	
	//Starting the getTicker goroutine
	go getTicker(getTickerChan)
	
	fmt.Println("Connecting to server on port 8000...")
	log.Fatalln(http.ListenAndServe(":8000", h)) //Set listening port (:8080). Handler is h indicating that chi router is used. log.Fatal checks for error and outputs if any.
}
