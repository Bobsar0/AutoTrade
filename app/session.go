//
package app

import(
	"github.com/chidi150c/autotrade/model"
	"github.com/chidi150c/autotrade/app/mock"
	"strings"
)
//Session carries the state parameters of a particular user
type Session struct{
	worker model.TransactionService
	GetTickerChan chan apiData
	GetBalanceChan chan apiData
}

//NewSession returns a new instance of *Session 
func NewSession () *Session{
	return &Session{
		GetTickerChan: make(chan apiData), //channel to communicate with getTicker goroutine
		GetBalanceChan: make(chan apiData),
	}
}

func (s *Session)SetWorker(host string){
	if strings.Contains(host, "mock") {
		s.worker = mock.NewWorker()
	}else if strings.Contains(host, "binance") {
		//s.worker = binance.NewWorker()
	}else if strings.Contains(host, "hitbtc") {
		//s.worker = hitbtc.NewWorker()
	}
}
