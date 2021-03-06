package todo

import (
	"context"
	"regexp"

	"todos/pkg/uuid"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

//Creator interface for mongodb transactions
type Creator interface {
	Insert(collection string, data Todo) error
}

//CreateLogic logical requirements for creating a Todo
type CreateLogic struct {
	DBManager Creator
	Log       logging.Logger
	UUID      uuid.UUIDGenerator
}

//CreateRequest request payload for creating a Todo
type CreateRequest struct {
	Status string `json:"status"`
	Item   string `json:"item"`
}

//Validate Validates the request
func (tdcl *CreateLogic) Validate(ctx context.Context, errors *ws.ServiceErrors, req *ws.Request) {
	tc := req.RequestBody.(*CreateRequest)
	TodoStatus := []string{"TODO", "DONE"}
	regex := `^[a-zA-Z][a-zA-Z0-9, ]*[._-]?[a-zA-Z0-9, ]+$`
	matched, _ := regexp.MatchString(regex, tc.Item)
	if !matched {
		errors.AddNewError(ws.Client, "INVALID_ITEM", "Item can contain only alphabets and numbers")
		errors.HTTPStatus = 422
	}
	if tc.Item == "" {
		errors.AddNewError(ws.Client, "EMPTY_ITEM", "Item is a required field")
		errors.HTTPStatus = 422
	}
	if tc.Status != "" && !contains(TodoStatus, tc.Status) {
		errors.AddNewError(ws.Client, "INVALID_STATUS", "Status should either be DONE or TODO")
		errors.HTTPStatus = 422
	}
}

//ProcessPayload Processes request
func (tdcl *CreateLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, tdcr *CreateRequest) {
	uuid, err := tdcl.UUID.Generate()
	if err != nil {
		tdcl.Log.LogErrorf("Error while generating UUID : Error %v", err)
		res.HTTPStatus = 500
		return
	}
	if tdcr.Status == "" {
		tdcr.Status = "TODO"
	}
	td := Todo{
		Item:   tdcr.Status,
		Status: tdcr.Item,
		TodoID: uuid,
	}
	tdcl.DBManager.Insert("todo", td)
	res.HTTPStatus = 201
	res.Body = map[string]string{
		"todoId": uuid,
	}
}
