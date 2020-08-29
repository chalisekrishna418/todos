package endpoint

import (
	"context"

	"todos/pkg/uuid"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

//TodoCreateLogic logical requirements for creating a Todo
type TodoCreateLogic struct {
	Log  logging.Logger
	UUID uuid.UUIDGenerator
}

//TodoCreateRequest request payload for creating a Todo
type TodoCreateRequest struct {
	Status string `json:"status"`
	Item   string `json:"item"`
}

//Validate Validates the request
func (tdcl *TodoCreateLogic) Validate(ctx context.Context, errors *ws.ServiceErrors, req *ws.Request) {
	tc := req.RequestBody.(*TodoCreateRequest)
	TodoStatus := []string{"TODO", "DONE"}
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
func (tdcl *TodoCreateLogic) ProcessPayload(ctx context.Context, errors *ws.ServiceErrors, req *ws.Request, res *ws.Response, tdcr *TodoCreateRequest) {
	uuid, err := tdcl.UUID.Generate()
	if err != nil {
		errors.AddNewError(ws.Client, "UUID_GENERATE_FAILED", "Error while generating TodoId")
		errors.HTTPStatus = 500
		// tdcl.Log.LogErrorf("Error while generating UUID : Error %v", err)
		return
	}
	if tdcr.Status == "" {
		tdcr.Status = "TODO"
	}
	res.HTTPStatus = 201
	res.Body = map[string]string{
		"TodoId": uuid,
	}
}
