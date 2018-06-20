package mock

import(
	"github.com/chidi150c/autotrade/model"
)

type Worker struct{

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

