//
package app

//Session carries the state parameters of a particular user
type Session struct{
	worker *Worker
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
