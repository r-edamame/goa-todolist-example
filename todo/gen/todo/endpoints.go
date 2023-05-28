// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo endpoints
//
// Command:
// $ goa gen todo/design

package todo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "todo" service endpoints.
type Endpoints struct {
	ListTasks    goa.Endpoint
	CreateTask   goa.Endpoint
	CompleteTask goa.Endpoint
	RevertTask   goa.Endpoint
}

// NewEndpoints wraps the methods of the "todo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		ListTasks:    NewListTasksEndpoint(s),
		CreateTask:   NewCreateTaskEndpoint(s),
		CompleteTask: NewCompleteTaskEndpoint(s),
		RevertTask:   NewRevertTaskEndpoint(s),
	}
}

// Use applies the given middleware to all the "todo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ListTasks = m(e.ListTasks)
	e.CreateTask = m(e.CreateTask)
	e.CompleteTask = m(e.CompleteTask)
	e.RevertTask = m(e.RevertTask)
}

// NewListTasksEndpoint returns an endpoint function that calls the method
// "listTasks" of service "todo".
func NewListTasksEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.ListTasks(ctx)
	}
}

// NewCreateTaskEndpoint returns an endpoint function that calls the method
// "createTask" of service "todo".
func NewCreateTaskEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateTaskPayload)
		return s.CreateTask(ctx, p)
	}
}

// NewCompleteTaskEndpoint returns an endpoint function that calls the method
// "completeTask" of service "todo".
func NewCompleteTaskEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CompleteTaskPayload)
		return s.CompleteTask(ctx, p)
	}
}

// NewRevertTaskEndpoint returns an endpoint function that calls the method
// "revertTask" of service "todo".
func NewRevertTaskEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*RevertTaskPayload)
		return s.RevertTask(ctx, p)
	}
}