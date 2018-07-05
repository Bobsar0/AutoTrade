//
package app

import (
	"log"
	"github.com/go-chi/chi" //Using chi mux/router
	"net/http"
	"github.com/bobsar0/AutoTrade/webClient"
	"github.com/bobsar0/AutoTrade/model"
)


var (
	indexTmpl = webClient.NewAppTemplate("index.gohtml")
	tickerTmpl = webClient.NewAppTemplate("getticker.gohtml")
	balanceTmpl = webClient.NewAppTemplate("getbalance.gohtml")
	placeorderTmpl = webClient.NewAppTemplate("placeorder.gohtml")
	orderFormTmpl = webClient.NewAppTemplate("orderform.gohtml") //Parse the order form

	OrderFormInput model.OrderInput
)

//AppHandler contains the chi mux and session and implements the http.ServeMux method, thus making it a handler
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
	h.mux.Get("/ticker", h.getTickerHandler) // getTickerHandler handles incoming requests from the trailing path '/ticker' and presents the ticker price to client/user 
	h.mux.Get("/balance", h.getBalanceHandler) // getBalanceHandler handles incoming requests from the trailing path '/balance' and presents the balance to client/user
	h.mux.Get("/placeorder", h.placeOrderHandler) // placeOrderHandler handles incoming requests from the path '/placeorder' and presents the output to client/user
	h.mux.Post("/placeorder", h.placeOrderHandler)

	h.mux.Get("/", h.indexHandler)
	return h
}


//AppHandler implements ServeHTTP method making it a Handler
func (h AppHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

//indexHandler delivers the Home page to the user
func (h *AppHandler)indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := indexTmpl.Execute(w,r,nil); err!=nil{
		log.Fatalln(err)
	}
}

//getTickerHandler presents the ticker price to the user
func (h *AppHandler)getTickerHandler(w http.ResponseWriter, r *http.Request){
	var host = "mock"
	h.session.SetWorker(host) //Sets the 'mock' worker for the user (app/mock/worker). Automatically sets it to other methods for the user as well since the receiver is a pointer
	tickerChan := make(chan interface{}) //tickerChan represents a channel that returns the ticker price
	h.session.GetTickerChan <- apiData{h.session.worker, tickerChan, OrderFormInput} //Send the content(ticker price) in tickerChan to retrieved from GetTicker(chan apiData) to user session
	ticker := <- tickerChan //ticker receives the ticker price via tickerChan
	if err := tickerTmpl.Execute(w, r, ticker); err!=nil{ //Execute template passing in 'ticker' as data to the template
		log.Fatalln(err)
	}
}

//getBalanceHandler presents the account balance to the user
func (h *AppHandler)getBalanceHandler(w http.ResponseWriter, r *http.Request){
	var host = "mock"
	h.session.SetWorker(host) //Sets the 'mock' worker for the user (app/mock/worker). Automatically sets it to other methods for the user as well since the receiver is a pointer
	balanceChan := make(chan interface{}) //balanceChan represents a channel that returns the balance
	h.session.GetBalanceChan <- apiData{h.session.worker, balanceChan, OrderFormInput} //Send the content(balance) in balanceChan retrieved from GetBalance(chan apiData) to user session
	balance := <- balanceChan //balance receives the account balance via balanceChan
	if err := balanceTmpl.Execute(w, r, balance); err!=nil{
		log.Fatalln(err)
	}
	return 
}

//placeOrderHandler presents output of placeOrder to the user
func (h *AppHandler)placeOrderHandler(w http.ResponseWriter, r *http.Request){
	var host = "mock"
	h.session.SetWorker(host) //Sets the 'mock' worker for the user (app/mock/worker). Automatically sets it to other methods for the user as well since the receiver is a pointer
	
	log.Println("placeOrderHandler method:", r.Method) //get request method
	
    if r.Method == "GET" { //if user requests for form
		if err := orderFormTmpl.Execute(w,r,nil); err!=nil{
			log.Fatalln(err)
		}
	} else { //if user submits form
		orderChan := make(chan interface{}) //orderChan represents a channel that returns the output of a placed order
		formInput := model.OrderInput{
			Symbol: r.FormValue("symbol"),
			Ticker: r.FormValue("price"),
			Quantity: r.FormValue("quantity"),
			Operation: r.FormValue("operation"),
		}
		h.session.PlaceOrderChan <- apiData{h.session.worker, orderChan, formInput} //Send the content(order) in placeorderChan to session
		orderOutput := <- orderChan //orderOutput receives the output of PlaceOrder() via placeOrderChan
		if err := placeorderTmpl.Execute(w, r, orderOutput); err!=nil{
			log.Fatalln(err)
		}
	return 
	}
}