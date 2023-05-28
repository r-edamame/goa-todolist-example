// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo service
//
// Command:
// $ goa gen todo/design

package todo

import (
	"context"
)

// todo service for management tasks
type Service interface {
	// ListTasks implements listTasks.
	ListTasks(context.Context) (res []*Task, err error)
	// CreateTask implements createTask.
	CreateTask(context.Context, *CreateTaskPayload) (res *Task, err error)
	// CompleteTask implements completeTask.
	CompleteTask(context.Context, *CompleteTaskPayload) (res *Task, err error)
	// RevertTask implements revertTask.
	RevertTask(context.Context, *RevertTaskPayload) (res *Task, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "todo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"listTasks", "createTask", "completeTask", "revertTask"}

// CompleteTaskPayload is the payload type of the todo service completeTask
// method.
type CompleteTaskPayload struct {
	// id of task
	ID *string
}

// CreateTaskPayload is the payload type of the todo service createTask method.
type CreateTaskPayload struct {
	// name of task
	Name *string
}

// RevertTaskPayload is the payload type of the todo service revertTask method.
type RevertTaskPayload struct {
	// id of task
	ID *string
}

// Task is the result type of the todo service createTask method.
type Task struct {
	// id of task
	ID *string
	// name of task
	Name *string
	// status of task
	Completed *bool
}