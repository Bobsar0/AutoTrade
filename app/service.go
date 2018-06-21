package app

import(
	"github.com/bobsar0/autotrade/model"
)

type apiData struct{
	w model.TransactionService
	callerChan chan interface{}
}

//PASSING CHANNEL-OVER-CHANNEL IMPLEMENTATION

//GetTicker() takes in a channel(gtc) which provides a channel of float64 (tickerChan)
//It provides the ticker price via a function that interacts with the API
//Then sends the ticker price back to the caller of the function (HANLDER) through the tickerChan
func GetTicker(gtc chan apiData){
	var ticker float64
	for{
		select{
		case aD := <- gtc: //if tickerChan receives a chan of float64 (gtc) from the getTicker() func caller
			ticker = aD.w.FuncThatReturnTicker() //Call the func to GET the ticker from the trading site (Using a test function for now)
			aD.callerChan <- ticker //Send the ticker back to the caller function(handler) via the tickerChan
		}
	}
}

func GetBalance(gtc chan apiData ){
	var acctBal float64
	for{
		select{
		case aD := <- gtc: //if balanceChan receives a chan of float64 (gtc) from the getBalance() func caller
			acctBal = aD.w.FuncThatReturnBalance() //Call the func to GET the ticker from the trading site (Using a test function for now)
			aD.callerChan <- acctBal //Send the balance back to the caller function(handler) via the balanceChan
		}
	}
}

func PlaceOrder(gtc chan apiData){
	var order model.OrderOutput
	for{
		select{
		case aD := <- gtc: //if tickerChan receives a chan of float64 (gtc) from the getTicker() func caller
			order = aD.w.FuncThatPlacesOrder() //Call the func to places order from the trading site (Using a test function for now)
			aD.callerChan <- order //Send the order output back to the caller function(handler) via the placeOrderChan
		}
	}
}