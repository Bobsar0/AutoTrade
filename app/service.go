package app

type apiData struct{
	w *Worker
	callerChan chan float64
}
//PASSING CHANNEL OVER CHANNEL
//getTicker() takes in a channel(gtc) which provides a channel of float64 (tickerChan)
//It provides the ticker price via a function that interacts with the API
//Then sends the ticker price back to the caller of the function (HANLDER) through the tickerChan
func GetTicker(gtc chan apiData){
	var ticker float64
	for{
		select{
		case td := <- gtc: //if tickerChan receives a chan of float64 (gtc) from the getTicker() func caller
			ticker = td.w.funcThatReturnTicker() //Call the func to GET the ticker from the trading site (Using a test function for now)
			td.callerChan <- ticker //Send the ticker back to the caller function(handler) via the tickerChan
		}
	}
}

func GetBalance(gtc chan apiData ){
	var acctBal float64
	for{
		select{
		case td:= <- gtc: //if tickerChan receives a chan of float64 (gtc) from the getTicker() func caller
			acctBal = td.w.funcThatReturnBalance() //Call the func to GET the ticker from the trading site (Using a test function for now)
			td.callerChan <- acctBal //Send the ticker back to the caller function(handler) via the tickerChan
		}
	}
}


type Worker struct{

}

func NewWorker() *Worker{
	return &Worker{
		
	}
}

//A test function to simulate getting ticker price from a trading site
//In the real world, this will be achieved by communicating with the site's API
func(w *Worker)funcThatReturnTicker()float64{
		return 0.002134442
	}
//A test function to simulate getting ticker price from a trading site
//In the real world, this will be achieved by communicating with the site's API
func(w *Worker)funcThatReturnBalance()float64{
	return 0.026654442
}

