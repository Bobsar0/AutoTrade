//
package app

import (
	"github.com/go-chi/chi" //Using chi mux/router
	"net/http"
	"fmt"
)

//AppHandler contains the chi mux and session and implements the http.ServeMux method, thus making it a handler
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
	h.mux.Get("/ticker", h.getTickerHandler) // getTickerHandler handles incoming requests from the trailing path '/ticker' and presents the ticker price to client/user 
	h.mux.Get("/balance", h.getBalanceHandler) // getBalanceHandler handles incoming requests from the trailing path '/balance' and presents the balance to client/user
	h.mux.Get("/placeorder", h.placeOrderHandler) // placeOrderHandler handles incoming requests from the path '/placeorder' and presents the output to client/user

	h.mux.Get("/", h.indexHandler)
	return h
}


//AppHandler implements ServeHTTP method making it a Handler
func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

//indexHandler delivers the Home page to the user
func (h *AppHandler)indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Welcome to AutoTrade<h1>
		<p><a href="/ticker">ticker</a></p>
		<p><a href="/balance">balance</a></p>
		<p><a href="/placeorder">Place Order</a></p>`) //Prints to webpage

}

//getTickerHandler presents the ticker price to the user
func (h *AppHandler)getTickerHandler(w http.ResponseWriter, r *http.Request){
	var host = "mock"
	h.session.SetWorker(host) //Sets the 'mock' worker for the user (app/mock/worker). Automatically sets it to other methods for the user as well since the receiver is a pointer
	tickerChan := make(chan interface{}) //tickerChan represents a channel that returns the ticker price
	h.session.GetTickerChan <- apiData{h.session.worker, tickerChan} //Send the content(ticker price) in tickerChan to retrieved from GetTicker(chan apiData) to user session
	ticker := <- tickerChan //ticker receives the ticker price via tickerChan
	responseToUser := fmt.Sprintf("<h1>Ticker: %.8f<h1>", ticker) //Returns response to user (which contains ticker) as a string
	fmt.Fprintf(w, "%s", responseToUser) //Prints response to user on the web page
	return 
}

//getBalanceHandler presents the account balance to the user
func (h *AppHandler)getBalanceHandler(w http.ResponseWriter, r *http.Request){
	balanceChan := make(chan interface{}) //balanceChan represents a channel that returns the balance
	h.session.GetBalanceChan <- apiData{h.session.worker, balanceChan} //Send the content(balance) in balanceChan retrieved from GetBalance(chan apiData) to user session
	balance := <- balanceChan //balance receives the account balance via balanceChan
	responseToUser := fmt.Sprintf("<h1>balance: %.8f<h1>", balance) //Returns response to user (which contains balance) as a string
	fmt.Fprintf(w, "%s", responseToUser) //Prints response to user on the web page
	return 
}

//getBalanceHandler presents output of placeOrder to the user
func (h *AppHandler)placeOrderHandler(w http.ResponseWriter, r *http.Request){
	placeOrderChan := make(chan interface{}) //balanceChan represents a channel that returns the balance
	h.session.PlaceOrderChan <- apiData{h.session.worker, placeOrderChan} //Send the content(balance) in balanceChan to session
	orderOutput := <- placeOrderChan //orderOutput receives the output of PlaceOrder() via placeOrderChan
	responseToUser := fmt.Sprintf(`<h1>Order Successful!</h1>
								<p>Order output:  %v </p>`, orderOutput) //Returns response to user (which contains placeOrder output) as a string
	fmt.Fprintf(w, "%s", responseToUser) //Prints response to user on the web page
	return 
}