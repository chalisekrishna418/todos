package todo

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/graniticio/granitic/test"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type mockLister struct {
	Data []interface{}
	Err  error
}

func (ml mockLister) List(string) ([]interface{}, error) {
	return ml.Data, ml.Err
}

func getTestTodoData() interface{} {
	return map[string]interface{}{
		"item":   "test 1",
		"status": "TODO",
		"todoId": "u-u-i-d",
	}
}

func TestList_Process(t *testing.T) {
	log := logging.CreateAnonymousLogger("TestLogger", logging.Fatal)
	t.Log("When request is successful returns Todo Items")
	{
		ll := ListLogic{
			DBManager: mockLister{
				Data: []interface{}{
					getTestTodoData(),
					getTestTodoData(),
				},
			},
			Log: log,
		}
		req := ws.Request{}
		res := ws.Response{}
		ctx := context.TODO()
		ll.Process(ctx, &req, &res)
		test.ExpectInt(t, res.HTTPStatus, 200)
		expjson := map[string]interface{}{
			"todos": []interface{}{
				getTestTodoData(),
				getTestTodoData(),
			},
		}
		if !reflect.DeepEqual(res.Body, expjson) {
			t.Fatalf("Expected %v, actual %v", expjson, res.Body)
		}
	}
	t.Log("When request is unsuccessful returns 500 status code")
	{
		ll := ListLogic{
			DBManager: mockLister{
				Data: []interface{}{
					getTestTodoData(),
					getTestTodoData(),
				},
				Err: fmt.Errorf("some error"),
			},
			Log: log,
		}
		req := ws.Request{}
		res := ws.Response{}
		ctx := context.TODO()
		ll.Process(ctx, &req, &res)
		test.ExpectInt(t, res.HTTPStatus, 500)
	}
}
