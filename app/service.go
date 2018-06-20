package app


//PASSING CHANNEL OVER CHANNEL
//getTicker() takes in a channel(gtc) which provides a channel of float64 (tickerChan)
//It provides the ticker price via a function that interacts with the API
//Then sends the ticker price back to the caller of the function (HANLDER) through the tickerChan
func GetTicker(gtc chan chan float64 ){
	var ticker float64
	for{
		select{
		case tickerChan := <- gtc: //if tickerChan receives a chan of float64 (gtc) from the getTicker() func caller
			ticker = functhatreturnticker() //Call the func to GET the ticker from the trading site (Using a test function for now)
			tickerChan <- ticker //Send the ticker back to the caller function(handler) via the tickerChan
		}
	}
}


//A test function to simulate getting ticker price from a trading site
//In the real world, this will be achieved by communicating with the site's API
func functhatreturnticker()float64{
	return 0.002134442
}

