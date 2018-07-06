//
package app

import(
	"github.com/bobsar0/AutoTrade/model"
	"github.com/bobsar0/AutoTrade/app/mock"
	"strings"
)
//Session carries the state parameters of a particular user
type Session struct{
	worker model.TransactionService
	GetTickerChan chan apiData
	GetBalanceChan chan apiData
	PlaceOrderChan chan apiData
	AddOrUpdateDbChan chan model.DbData 
	GetDbChan chan model.DbData
}

//NewSession returns a new instance of *Session 
func NewSession () *Session{
	return &Session{
		GetTickerChan: make(chan apiData), //channel to communicate with getTicker goroutine
		GetBalanceChan: make(chan apiData),//channel to communicate with getBalance goroutine
		PlaceOrderChan: make(chan apiData),//channel to communicate with placeOrder goroutine
	}
}

var _ model.Session = &Session{}

//Authenticate authorises user to access his account
func (s *Session)Authenticate()*model.User{
	return &model.User{
		Username: "Steve",
	}
}

//SetWorker sets the worker for the user based on the trading site selected
func (s *Session)SetWorker(host string)error{
	if s == nil {
		return 	model.ErrNilSessionStruct		
	}
	if strings.Contains(host, "mock") {
		wkr := mock.NewWorker()
		wkr.Sess = s
		//wkr.GetTickerChan = s.GetTickerChan
		//wkr.GetBalanceChan = s.GetBalanceChan
		wkr.AddOrUpdateDbChan = s.AddOrUpdateDbChan
		wkr.GetDbChan = s.GetDbChan
		//wkr.DeleteDbChan = s.DeleteDbChan
		s.worker = wkr
		return nil
	}
	if strings.Contains(host, "exchPlatform1") {
		//s.worker = exchPlatform1.NewWorker()
		return nil
	}
	if strings.Contains(host, "exchPlatform2") {
		//s.worker = exchPlatform2.NewWorker()
		return nil
	}
	return nil
}
