//
package app

import (
	"github.com/go-chi/chi" //Using chi mux/router
	"net/http"
	"fmt"
	
)

//Type AppHandler contains the chi mux and session and implements the ServeMux method
type AppHandler struct{
	mux *chi.Mux
	session *Session
}


//NewAppHandler returns a new instance of *AppHandler
func NewAppHandler (s *Session) AppHandler{
	h := AppHandler{
		mux: chi.NewRouter(),
		session: s,
	}
	h.mux.Get("/ticker", h.getTickerHandler)
	h.mux.Get("/balance", h.getBalanceHandler)
	h.mux.Get("/", h.indexHandler)
	return h
}


//AppHandler implements ServeHTTP method making it a Handler
func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

//indexHandler delivers the Home page to the user
func (h *AppHandler)indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Welcome to AutoTrade<h1>
		<p><a href="/ticker">ticker</a></p>
		<p><a href="/balance">balance</a></p>`) //Prints ticker to webpage
}

//getTickerHandler presents the ticker price to the user
func (h *AppHandler)getTickerHandler(w http.ResponseWriter, r *http.Request){
	var host = "mock"
	h.session.SetWorker(host)
	tickerChan := make(chan float64) //tickerChan represents a channel that returns the ticker price
	h.session.GetTickerChan <- apiData{h.session.worker, tickerChan} //Send the content(ticker price) in tickerChan to session
	ticker := <- tickerChan //ticker receives the ticker price via tickerChan
	responseToUser := fmt.Sprintf("<h1>Ticker: %.8f<h1>", ticker) //Returns response to user (which contains ticker) as a string
	fmt.Fprintf(w, "%s", responseToUser) //Prints response to user on the web page
	return 
}

//getBalanceHandler presents the account balance to the user
func (h *AppHandler)getBalanceHandler(w http.ResponseWriter, r *http.Request){
	balanceChan := make(chan float64) //balanceChan represents a channel that returns the balance
	h.session.GetBalanceChan <- apiData{h.session.worker, balanceChan} //Send the content(balance) in balanceChan to session
	balance := <- balanceChan //balance receives the account balance via balanceChan
	responseToUser := fmt.Sprintf("<h1>balance: %.8f<h1>", balance) //Returns response to user (which contains balance) as a string
	fmt.Fprintf(w, "%s", responseToUser) //Prints response to user on the web page
	return 
}