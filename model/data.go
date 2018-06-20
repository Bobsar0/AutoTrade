package model

type TransactionID string

type Transaction struct{
	ID TransactionID
	Order float64
	Price float64
	tYPE  string
	Operation string
}

type TransactionService interface{
	FuncThatReturnTicker() float64
	FuncThatReturnBalance()float64
}