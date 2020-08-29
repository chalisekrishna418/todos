package endpoint

import (
	"context"
	"testing"

	"github.com/graniticio/granitic/test"
	"github.com/graniticio/granitic/v2/ws"
)

func getTestTodoCreateLogic() TodoCreateLogic {
	return TodoCreateLogic{
		TodoID: "a-b-c-d",
	}
}

func getTestTodoCreateRequest() TodoCreateRequest {
	return TodoCreateRequest{
		Item:   "Task 1",
		Status: "TODO",
	}
}

func TestCreate_validate(t *testing.T) {
	t.Log("when Item is empty")
	{
		tdcl := getTestTodoCreateLogic()
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
	}
	t.Log("when Item is invalid")
	{
		tdcl := getTestTodoCreateLogic()
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
	}
	t.Log("When status is nil")
	{
		tdcl := getTestTodoCreateLogic()
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
