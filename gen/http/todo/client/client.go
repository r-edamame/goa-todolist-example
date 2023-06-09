// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo client HTTP transport
//
// Command:
// $ goa gen todo/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the todo service endpoint HTTP clients.
type Client struct {
	// ListTasks Doer is the HTTP client used to make requests to the listTasks
	// endpoint.
	ListTasksDoer goahttp.Doer

	// CreateTask Doer is the HTTP client used to make requests to the createTask
	// endpoint.
	CreateTaskDoer goahttp.Doer

	// CompleteTask Doer is the HTTP client used to make requests to the
	// completeTask endpoint.
	CompleteTaskDoer goahttp.Doer

	// RevertTask Doer is the HTTP client used to make requests to the revertTask
	// endpoint.
	RevertTaskDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the todo service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListTasksDoer:       doer,
		CreateTaskDoer:      doer,
		CompleteTaskDoer:    doer,
		RevertTaskDoer:      doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// ListTasks returns an endpoint that makes HTTP requests to the todo service
// listTasks server.
func (c *Client) ListTasks() goa.Endpoint {
	var (
		decodeResponse = DecodeListTasksResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildListTasksRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListTasksDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("todo", "listTasks", err)
		}
		return decodeResponse(resp)
	}
}

// CreateTask returns an endpoint that makes HTTP requests to the todo service
// createTask server.
func (c *Client) CreateTask() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateTaskRequest(c.encoder)
		decodeResponse = DecodeCreateTaskResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCreateTaskRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateTaskDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("todo", "createTask", err)
		}
		return decodeResponse(resp)
	}
}

// CompleteTask returns an endpoint that makes HTTP requests to the todo
// service completeTask server.
func (c *Client) CompleteTask() goa.Endpoint {
	var (
		decodeResponse = DecodeCompleteTaskResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCompleteTaskRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CompleteTaskDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("todo", "completeTask", err)
		}
		return decodeResponse(resp)
	}
}

// RevertTask returns an endpoint that makes HTTP requests to the todo service
// revertTask server.
func (c *Client) RevertTask() goa.Endpoint {
	var (
		decodeResponse = DecodeRevertTaskResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRevertTaskRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RevertTaskDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("todo", "revertTask", err)
		}
		return decodeResponse(resp)
	}
}
