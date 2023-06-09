// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo gRPC client encoders and decoders
//
// Command:
// $ goa gen todo/design

package client

import (
	"context"
	todopb "todo/gen/grpc/todo/pb"
	todo "todo/gen/todo"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildCreateTaskFunc builds the remote method to invoke for "todo" service
// "createTask" endpoint.
func BuildCreateTaskFunc(grpccli todopb.TodoClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.CreateTask(ctx, reqpb.(*todopb.CreateTaskRequest), opts...)
		}
		return grpccli.CreateTask(ctx, &todopb.CreateTaskRequest{}, opts...)
	}
}

// EncodeCreateTaskRequest encodes requests sent to todo createTask endpoint.
func EncodeCreateTaskRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*todo.CreateTaskPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("todo", "createTask", "*todo.CreateTaskPayload", v)
	}
	return NewProtoCreateTaskRequest(payload), nil
}

// DecodeCreateTaskResponse decodes responses from the todo createTask endpoint.
func DecodeCreateTaskResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*todopb.CreateTaskResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("todo", "createTask", "*todopb.CreateTaskResponse", v)
	}
	res := NewCreateTaskResult(message)
	return res, nil
}

// BuildCompleteTaskFunc builds the remote method to invoke for "todo" service
// "completeTask" endpoint.
func BuildCompleteTaskFunc(grpccli todopb.TodoClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.CompleteTask(ctx, reqpb.(*todopb.CompleteTaskRequest), opts...)
		}
		return grpccli.CompleteTask(ctx, &todopb.CompleteTaskRequest{}, opts...)
	}
}

// EncodeCompleteTaskRequest encodes requests sent to todo completeTask
// endpoint.
func EncodeCompleteTaskRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*todo.CompleteTaskPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("todo", "completeTask", "*todo.CompleteTaskPayload", v)
	}
	return NewProtoCompleteTaskRequest(payload), nil
}

// DecodeCompleteTaskResponse decodes responses from the todo completeTask
// endpoint.
func DecodeCompleteTaskResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*todopb.CompleteTaskResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("todo", "completeTask", "*todopb.CompleteTaskResponse", v)
	}
	res := NewCompleteTaskResult(message)
	return res, nil
}

// BuildRevertTaskFunc builds the remote method to invoke for "todo" service
// "revertTask" endpoint.
func BuildRevertTaskFunc(grpccli todopb.TodoClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.RevertTask(ctx, reqpb.(*todopb.RevertTaskRequest), opts...)
		}
		return grpccli.RevertTask(ctx, &todopb.RevertTaskRequest{}, opts...)
	}
}

// EncodeRevertTaskRequest encodes requests sent to todo revertTask endpoint.
func EncodeRevertTaskRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*todo.RevertTaskPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("todo", "revertTask", "*todo.RevertTaskPayload", v)
	}
	return NewProtoRevertTaskRequest(payload), nil
}

// DecodeRevertTaskResponse decodes responses from the todo revertTask endpoint.
func DecodeRevertTaskResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*todopb.RevertTaskResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("todo", "revertTask", "*todopb.RevertTaskResponse", v)
	}
	res := NewRevertTaskResult(message)
	return res, nil
}
