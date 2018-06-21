//Package model stores data needed for transactions
package model

type TransactionID string

//Transaction type contains data needed for a buy/sell transaction
type Transaction struct{
	ID TransactionID
	Order float64 //Quantity
	Price float64
	tYPE  string
	Operation string //Buy/Sell
}

//TransactionService interface contains all the methods needed to interact with the trading site's API
//Worker type implements these methods and thus implements the interface
type TransactionService interface{
	FuncThatReturnTicker() float64 //Test function that returns ticker price
	FuncThatReturnBalance()float64 //Test function that returns balance
	FuncThatPlacesOrder() OrderOutput// Test function that places order
}

//OrderOutput fields contains info about an order after it has been successfully placed
type OrderOutput struct{
	OrderID int
	ClientID int
	Symbol string
	Ticker float64
	Quantity float64
	Balance float64
}