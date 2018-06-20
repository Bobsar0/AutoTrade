//
package app

//Session carries the state parameters of a particular user
type Session struct{
	GetTickerChan chan chan float64
}

//NewSession returns a new instance of *Session 
func NewSession () *Session{
	return &Session{
		GetTickerChan: make(chan chan float64), //channel to communicate with getTicker goroutine
	}
}
