package todo

import (
	"context"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

//Lister interface for mongodb list transactions
type Lister interface {
	List(collection string) ([]interface{}, error)
}

//ListLogic requirements for listing
type ListLogic struct {
	DBManager Lister
	Log       logging.Logger
}

//Process Processes the request
func (ll *ListLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	tds, err := ll.DBManager.List("todo")
	if err != nil {
		ll.Log.LogErrorf("Error: %v", err)
		res.HTTPStatus = 500
		return
	}
	res.HTTPStatus = 200
	res.Body = map[string]interface{}{
		"todos": tds,
	}
}
