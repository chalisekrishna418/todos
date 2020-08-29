package endpoint

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/graniticio/granitic/test"
	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type mockUUIDGenerator struct {
	UUID string
	Err  error
}

func (mg mockUUIDGenerator) Generate() (string, error) {
	return mg.UUID, mg.Err
}

func getTestTodoCreateRequest() TodoCreateRequest {
	return TodoCreateRequest{
		Item:   "Task 1",
		Status: "TODO",
	}
}

func TestCreate_Validate(t *testing.T) {
	log := logging.CreateAnonymousLogger("TestLogger", logging.Fatal)
	t.Log("when Item is empty")
	{
		tdcl := TodoCreateLogic{
			Log: log,
			UUID: mockUUIDGenerator{
				UUID: "u-u-i-d",
			},
		}
		tdcr := getTestTodoCreateRequest()
		tdcr.Item = ""
		req := ws.Request{}
		se := ws.ServiceErrors{}
		req.RequestBody = &tdcr
		ctx := context.TODO()
		tdcl.Validate(ctx, &se, &req)
		errs := se.Errors
		test.ExpectInt(t, len(errs), 1)
		test.ExpectString(t, errs[0].Message, "Item is a required field")
		test.ExpectString(t, errs[0].Code, "EMPTY_ITEM")
		test.ExpectInt(t, se.HTTPStatus, 422)
	}
	t.Log("when Item is invalid")
	{
		tdcl := TodoCreateLogic{
			Log: log,
			UUID: mockUUIDGenerator{
				UUID: "u-u-i-d",
			},
		}
		tdcr := getTestTodoCreateRequest()
		tdcr.Status = "xyz"
		req := ws.Request{}
		se := ws.ServiceErrors{}
		req.RequestBody = &tdcr
		ctx := context.TODO()
		tdcl.Validate(ctx, &se, &req)
		errs := se.Errors
		test.ExpectInt(t, len(errs), 1)
		test.ExpectString(t, errs[0].Message, "Status should either be DONE or TODO")
		test.ExpectString(t, errs[0].Code, "INVALID_STATUS")
		test.ExpectInt(t, se.HTTPStatus, 422)
	}
	t.Log("When status is nil")
	{
		tdcl := TodoCreateLogic{
			Log:  log,
			UUID: mockUUIDGenerator{UUID: "u-u-i-d"},
		}
		tdcr := getTestTodoCreateRequest()
		tdcr.Status = ""
		req := ws.Request{}
		se := ws.ServiceErrors{}
		req.RequestBody = &tdcr
		ctx := context.TODO()
		tdcl.Validate(ctx, &se, &req)
		errs := se.Errors
		test.ExpectInt(t, len(errs), 0)
	}
}

func TestCreate_ProcessPayload(t *testing.T) {
	log := logging.CreateAnonymousLogger("TestLogger", logging.Fatal)
	t.Log("When request is successful, returns uuid")
	{
		tdcl := TodoCreateLogic{
			Log: log,
			UUID: mockUUIDGenerator{
				UUID: "u-u-i-d",
			},
		}
		tdcr := getTestTodoCreateRequest()
		req := ws.Request{}
		res := ws.Response{}
		ctx := context.TODO()
		se := ws.ServiceErrors{}
		tdcl.ProcessPayload(ctx, &se, &req, &res, &tdcr)
		test.ExpectInt(t, res.HTTPStatus, 201)
		expBody := map[string]string{
			"TodoId": "u-u-i-d",
		}
		if !reflect.DeepEqual(res.Body, expBody) {
			t.Fatalf("Expected %v, actual %v", expBody, res.Body)
		}
	}
	t.Log("When status is empty: request is successful, returns uuid")
	{
		tdcl := TodoCreateLogic{
			Log: log,
			UUID: mockUUIDGenerator{
				UUID: "u-u-i-d",
			},
		}
		tdcr := getTestTodoCreateRequest()
		tdcr.Status = ""
		req := ws.Request{}
		res := ws.Response{}
		ctx := context.TODO()
		se := ws.ServiceErrors{}
		tdcl.ProcessPayload(ctx, &se, &req, &res, &tdcr)
		test.ExpectInt(t, res.HTTPStatus, 201)
		expBody := map[string]string{
			"TodoId": "u-u-i-d",
		}
		if !reflect.DeepEqual(res.Body, expBody) {
			t.Fatalf("Expected %v, actual %v", expBody, res.Body)
		}
	}
	t.Log("When request is unsuccessful, returns error")
	{
		tdcl := TodoCreateLogic{
			Log: log,
			UUID: mockUUIDGenerator{
				Err: fmt.Errorf("some error"),
			},
		}
		tdcr := getTestTodoCreateRequest()
		req := ws.Request{}
		res := ws.Response{}
		ctx := context.TODO()
		se := ws.ServiceErrors{}
		tdcl.ProcessPayload(ctx, &se, &req, &res, &tdcr)
		test.ExpectInt(t, se.HTTPStatus, 500)
		test.ExpectInt(t, len(se.Errors), 1)
		test.ExpectString(t, se.Errors[0].Message, "Error while generating TodoId")
		test.ExpectString(t, se.Errors[0].Code, "UUID_GENERATE_FAILED")
	}
}
