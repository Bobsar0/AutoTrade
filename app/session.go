//
package app

import(
	"github.com/bobsar0/autotrade/model"
	"github.com/bobsar0/autotrade/app/mock"
	"strings"
)
//Session carries the state parameters of a particular user
type Session struct{
	worker model.TransactionService
	GetTickerChan chan apiData
	GetBalanceChan chan apiData
	PlaceOrderChan chan apiData
}

//NewSession returns a new instance of *Session 
func NewSession () *Session{
	return &Session{
		GetTickerChan: make(chan apiData), //channel to communicate with getTicker goroutine
		GetBalanceChan: make(chan apiData),//channel to communicate with getBalance goroutine
		PlaceOrderChan: make(chan apiData),//channel to communicate with placeOrder goroutine
	}
}

//SetWorker sets the worker for the user based on the trading site selected
func (s *Session)SetWorker(host string){
	if strings.Contains(host, "mock") {
		s.worker = mock.NewWorker()
	}else if strings.Contains(host, "binance") {
		//s.worker = binance.NewWorker()
	}else if strings.Contains(host, "hitbtc") {
		//s.worker = hitbtc.NewWorker()
	}
}
