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
func NewAppHandler (s *Session) *AppHandler{
	h := &AppHandler{
		mux: chi.NewRouter(),
		session: s,
	}
	h.mux.Get("/ticker", h.getTickerHandler)
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
		<p><a href="/ticker">ticker</a></p>`) //Prints ticker to webpage
}

//getTickerHandler presents the ticker price to the user
func (h *AppHandler)getTickerHandler(w http.ResponseWriter, r *http.Request){
	tickerChan := make(chan float64) //tickerChan represents a channel that returns the ticker price
	h.session.GetTickerChan <- tickerChan //Send the content(ticker price) in tickerChan to session
	ticker := <- tickerChan //ticker receives the ticker price via tickerChan
	
	responseToUser := fmt.Sprintf("<h1>Ticker: %.8f<h1>", ticker) //Returns response to user (which contains ticker) as a string
	fmt.Fprintf(w, "%s", responseToUser) //Prints response to user on the web page
	return 
}