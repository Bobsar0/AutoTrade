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
	FuncThatPlacesOrder(OrderInput) OrderOutput// Test function that places order
}

//Mock inputs needed to place an order
type OrderInput struct{
	Symbol string
	Quantity string
	Ticker string
	Operation string
}

//Mock output of a successful placing of an order
type OrderOutput struct{
	OrderID int
	ClientID int
	Symbol string
	Ticker string
	Quantity float64
	Balance float64
}