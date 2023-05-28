// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo HTTP client types
//
// Command:
// $ goa gen todo/design

package client

import (
	todo "todo/gen/todo"
)

// CreateTaskRequestBody is the type of the "todo" service "createTask"
// endpoint HTTP request body.
type CreateTaskRequestBody struct {
	// name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ListTasksResponseBody is the type of the "todo" service "listTasks" endpoint
// HTTP response body.
type ListTasksResponseBody []*TaskResponse

// CreateTaskResponseBody is the type of the "todo" service "createTask"
// endpoint HTTP response body.
type CreateTaskResponseBody struct {
	// id of task
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// status of task
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// CompleteTaskResponseBody is the type of the "todo" service "completeTask"
// endpoint HTTP response body.
type CompleteTaskResponseBody struct {
	// id of task
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// status of task
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// RevertTaskResponseBody is the type of the "todo" service "revertTask"
// endpoint HTTP response body.
type RevertTaskResponseBody struct {
	// id of task
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// status of task
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// TaskResponse is used to define fields on response body types.
type TaskResponse struct {
	// id of task
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// status of task
	Completed *bool `form:"completed,omitempty" json:"completed,omitempty" xml:"completed,omitempty"`
}

// NewCreateTaskRequestBody builds the HTTP request body from the payload of
// the "createTask" endpoint of the "todo" service.
func NewCreateTaskRequestBody(p *todo.CreateTaskPayload) *CreateTaskRequestBody {
	body := &CreateTaskRequestBody{
		Name: p.Name,
	}
	return body
}

// NewListTasksTaskOK builds a "todo" service "listTasks" endpoint result from
// a HTTP "OK" response.
func NewListTasksTaskOK(body []*TaskResponse) []*todo.Task {
	v := make([]*todo.Task, len(body))
	for i, val := range body {
		v[i] = unmarshalTaskResponseToTodoTask(val)
	}

	return v
}

// NewCreateTaskTaskOK builds a "todo" service "createTask" endpoint result
// from a HTTP "OK" response.
func NewCreateTaskTaskOK(body *CreateTaskResponseBody) *todo.Task {
	v := &todo.Task{
		ID:        body.ID,
		Name:      body.Name,
		Completed: body.Completed,
	}

	return v
}

// NewCompleteTaskTaskOK builds a "todo" service "completeTask" endpoint result
// from a HTTP "OK" response.
func NewCompleteTaskTaskOK(body *CompleteTaskResponseBody) *todo.Task {
	v := &todo.Task{
		ID:        body.ID,
		Name:      body.Name,
		Completed: body.Completed,
	}

	return v
}

// NewRevertTaskTaskOK builds a "todo" service "revertTask" endpoint result
// from a HTTP "OK" response.
func NewRevertTaskTaskOK(body *RevertTaskResponseBody) *todo.Task {
	v := &todo.Task{
		ID:        body.ID,
		Name:      body.Name,
		Completed: body.Completed,
	}

	return v
}