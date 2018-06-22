package model

type TransactionID string

type Transaction struct{
	ID TransactionID
	Order float64
	Price float64
	tYPE  string
	Operation string
	Token string
}

type TransactionService interface{
	FuncThatReturnTicker() float64 
	FuncThatReturnBalance()float64
	AddTransaction(Transaction) (TransactionID, error)
	GetTransaction(TransactionID) (Transaction, error)
}

type UserID string

type User struct{
	ID UserID
	Username string
	Password string
	PublicKey string
	Secret string
	Token string
}

type UserService interface{
	AddUser(*User) (UserID, error)
	GetUser(UserID) (*User, error) 
	UpdateUser(*User) error
	ListUsers() ([]*User, error) 
	DeleteUser (UserID) error 
}

type DbData struct{
	TransID TransactionID
	Transaction Transaction
	CallerChan chan DbResp
}

type DbResp struct{
	TransID TransactionID
	Transaction Transaction
	Err error 
}

type ApiData struct{
	W TransactionService
	CallerChan chan float64
}

type Session interface{
	Authenticate()*User
	SetWorker(string)error
}