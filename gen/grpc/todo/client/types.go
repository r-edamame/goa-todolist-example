// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo gRPC client types
//
// Command:
// $ goa gen todo/design

package client

import (
	todopb "todo/gen/grpc/todo/pb"
	todo "todo/gen/todo"
)

// NewProtoCreateTaskRequest builds the gRPC request type from the payload of
// the "createTask" endpoint of the "todo" service.
func NewProtoCreateTaskRequest(payload *todo.CreateTaskPayload) *todopb.CreateTaskRequest {
	message := &todopb.CreateTaskRequest{
		Name: payload.Name,
	}
	return message
}

// NewCreateTaskResult builds the result type of the "createTask" endpoint of
// the "todo" service from the gRPC response type.
func NewCreateTaskResult(message *todopb.CreateTaskResponse) *todo.Task {
	result := &todo.Task{
		ID:        message.Id,
		Name:      message.Name,
		Completed: message.Completed,
	}
	return result
}

// NewProtoCompleteTaskRequest builds the gRPC request type from the payload of
// the "completeTask" endpoint of the "todo" service.
func NewProtoCompleteTaskRequest(payload *todo.CompleteTaskPayload) *todopb.CompleteTaskRequest {
	message := &todopb.CompleteTaskRequest{
		Id: payload.ID,
	}
	return message
}

// NewCompleteTaskResult builds the result type of the "completeTask" endpoint
// of the "todo" service from the gRPC response type.
func NewCompleteTaskResult(message *todopb.CompleteTaskResponse) *todo.Task {
	result := &todo.Task{
		ID:        message.Id,
		Name:      message.Name,
		Completed: message.Completed,
	}
	return result
}

// NewProtoRevertTaskRequest builds the gRPC request type from the payload of
// the "revertTask" endpoint of the "todo" service.
func NewProtoRevertTaskRequest(payload *todo.RevertTaskPayload) *todopb.RevertTaskRequest {
	message := &todopb.RevertTaskRequest{
		Id: payload.ID,
	}
	return message
}

// NewRevertTaskResult builds the result type of the "revertTask" endpoint of
// the "todo" service from the gRPC response type.
func NewRevertTaskResult(message *todopb.RevertTaskResponse) *todo.Task {
	result := &todo.Task{
		ID:        message.Id,
		Name:      message.Name,
		Completed: message.Completed,
	}
	return result
}
