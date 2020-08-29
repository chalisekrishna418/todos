package endpoint

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/graniticio/granitic/v2/ws"
)

//TodoCreateLogic logical requirements for creating a Todo
type TodoCreateLogic struct {
	TodoID string
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
func (tdcl *TodoCreateLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, tdcr *TodoCreateRequest) {
	uuid, err := uuid.NewRandom()
	if tdcr.Status == "" {
		tdcr.Status = "TODO"
	}
	if err != nil {
		fmt.Errorf("[Error] Generating uuid")
	}
	res.Body = map[string]string{
		"TodoId": uuid.String(),
	}
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
