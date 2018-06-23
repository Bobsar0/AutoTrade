package mock

import(
	"github.com/chidi150c/autotrade/model"
	//"log"
	//"github.com/pkg/errors"
)

type Worker struct{
	AddOrUpdateDbChan chan model.DbData
	GetDbChan chan model.DbData
}

func NewWorker() *Worker{
	return &Worker{
		
	}
}

var _ model.TransactionService = &Worker{}
//A test function to simulate getting ticker price from a trading site
//In the real world, this will be achieved by communicating with the site's API
func(w *Worker)FuncThatReturnTicker()float64{
		return 0.002134442
	}
//A test function to simulate getting ticker price from a trading site
//In the real world, this will be achieved by communicating with the site's API
func(w *Worker)FuncThatReturnBalance()float64{
	return 0.026654442
}

func(w *Worker)AddTransaction(ts model.Transaction) (model.TransactionID, error){
	cChan := make(chan model.DbResp)
	w.AddOrUpdateDbChan <-model.DbData{Transaction: ts, CallerChan: cChan}
	res := <-cChan
	return res.TransID, res.Err
}

func(w *Worker)GetTransaction(id model.TransactionID) (model.Transaction, error){
	cChan := make(chan model.DbResp)
	w.GetDbChan <-model.DbData{TransID: id, CallerChan: cChan}
	res := <-cChan
	return res.Transaction, res.Err
}
