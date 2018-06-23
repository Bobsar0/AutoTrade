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
	AddOrUpdateDbChan chan model.DbData 
	GetDbChan chan model.DbData
	DeleteDbChan chan model.DbData
}

//NewSession returns a new instance of *Session 
func NewSession () *Session{
	return &Session{
		GetTickerChan: make(chan apiData), //channel to communicate with getTicker goroutine
		GetBalanceChan: make(chan apiData),
		AddOrUpdateDbChan: make(chan model.DbData),
		GetDbChan: make(chan model.DbData),
		DeleteDbChan: make(chan model.DbData),
	}
}

//var _ model.Session = &Session{}

func (s *Session)SetWorker(host string)error{
	if s == nil {
		return 	model.ErrNilSessionStruct		
	}
	if strings.Contains(host, "mock") {
		wkr := mock.NewWorker()
		//wkr.Sess = s
		//wkr.GetTickerChan = s.GetTickerChan
		//wkr.GetBalanceChan = s.GetBalanceChan
		wkr.AddOrUpdateDbChan = s.AddOrUpdateDbChan
		wkr.GetDbChan = s.GetDbChan
		//wkr.DeleteDbChan = s.DeleteDbChan
		s.worker = wkr
		return nil
	}
	if strings.Contains(host, "binance") {
		//s.worker = binance.NewWorker()
		return nil
	}
	if strings.Contains(host, "hitbtc") {
		//s.worker = hitbtc.NewWorker()
		return nil
	}
	return nil
}
