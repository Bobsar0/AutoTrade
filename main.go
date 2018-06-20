//Using http package to build a simple server
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi"
)

func (h *appHandler)indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Welcome to AutoTrade<h1>
		<p><a href="/ticker">ticker</a></p>`) //Prints to webpage
}
func (h *appHandler)getTickerHandler(w http.ResponseWriter, r *http.Request){
	tickerChan := make(chan float64)
	h.session.getTickerChan <- tickerChan 
	ticker := <- tickerChan 
	responseToUser := fmt.Sprintf("<h1>Ticker: %.8f<h1>", ticker)
	fmt.Fprintf(w, "%s", responseToUser)
	return 
}

//A test function to simulate getting ticker price from a trading site
func functhatreturnticker()float64{
	return 0.002134442
}

func getTicker(gtc chan chan float64 ){
	var ticker float64
	for{
		select{
		case tickerChan := <- gtc:
			ticker = functhatreturnticker()
			tickerChan <- ticker
		}
	}
}

type Session struct{
	getTickerChan chan chan float64
}

func NewSession () *Session{
	return &Session{

	}
}

type appHandler struct{
	mux *chi.Mux
	session *Session
}

func NewAppHandler (s *Session) *appHandler{
	h := &appHandler{
		mux: chi.NewRouter(),
		session: s,
	}
	h.mux.Get("/ticker", h.getTickerHandler)
	h.mux.Get("/", h.indexHandler)
	return h
}

func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}


func main() {
	//channel to communicate with getTicker goroutine
	getTickerChan := make(chan chan float64)
	s := NewSession()
	h := NewAppHandler (s)
	h.session.getTickerChan = getTickerChan
	//Starting the getTicker goroutine
	go getTicker(getTickerChan)
	
	fmt.Println("Connecting to server on port 8000...")
	log.Fatalln(http.ListenAndServe(":8000", h)) //Set listening port (:8080). Handler is nil indicating that DefaultServeMux should be used. log.Fatal checks for error and outputs if any.
}
