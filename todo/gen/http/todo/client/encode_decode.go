// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo HTTP client encoders and decoders
//
// Command:
// $ goa gen todo/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	todo "todo/gen/todo"

	goahttp "goa.design/goa/v3/http"
)

// BuildListTasksRequest instantiates a HTTP request object with method and
// path set to call the "todo" service "listTasks" endpoint
func (c *Client) BuildListTasksRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListTasksTodoPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "listTasks", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListTasksResponse returns a decoder for responses returned by the todo
// listTasks endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListTasksResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListTasksResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "listTasks", err)
			}
			res := NewListTasksTaskOK(body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "listTasks", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateTaskRequest instantiates a HTTP request object with method and
// path set to call the "todo" service "createTask" endpoint
func (c *Client) BuildCreateTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateTaskTodoPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "createTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateTaskRequest returns an encoder for requests sent to the todo
// createTask server.
func EncodeCreateTaskRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*todo.CreateTaskPayload)
		if !ok {
			return goahttp.ErrInvalidType("todo", "createTask", "*todo.CreateTaskPayload", v)
		}
		body := NewCreateTaskRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("todo", "createTask", err)
		}
		return nil
	}
}

// DecodeCreateTaskResponse returns a decoder for responses returned by the
// todo createTask endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateTaskResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateTaskResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "createTask", err)
			}
			res := NewCreateTaskTaskOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "createTask", resp.StatusCode, string(body))
		}
	}
}

// BuildCompleteTaskRequest instantiates a HTTP request object with method and
// path set to call the "todo" service "completeTask" endpoint
func (c *Client) BuildCompleteTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*todo.CompleteTaskPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "completeTask", "*todo.CompleteTaskPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CompleteTaskTodoPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "completeTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeCompleteTaskResponse returns a decoder for responses returned by the
// todo completeTask endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCompleteTaskResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CompleteTaskResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "completeTask", err)
			}
			res := NewCompleteTaskTaskOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "completeTask", resp.StatusCode, string(body))
		}
	}
}

// BuildRevertTaskRequest instantiates a HTTP request object with method and
// path set to call the "todo" service "revertTask" endpoint
func (c *Client) BuildRevertTaskRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*todo.RevertTaskPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("todo", "revertTask", "*todo.RevertTaskPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RevertTaskTodoPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "revertTask", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRevertTaskResponse returns a decoder for responses returned by the
// todo revertTask endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeRevertTaskResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RevertTaskResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "revertTask", err)
			}
			res := NewRevertTaskTaskOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "revertTask", resp.StatusCode, string(body))
		}
	}
}

// unmarshalTaskResponseToTodoTask builds a value of type *todo.Task from a
// value of type *TaskResponse.
func unmarshalTaskResponseToTodoTask(v *TaskResponse) *todo.Task {
	res := &todo.Task{
		ID:        v.ID,
		Name:      v.Name,
		Completed: v.Completed,
	}

	return res
}
