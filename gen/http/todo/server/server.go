// Code generated by goa v3.11.3, DO NOT EDIT.
//
// todo HTTP server
//
// Command:
// $ goa gen todo/design

package server

import (
	"context"
	"net/http"
	"regexp"
	todo "todo/gen/todo"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the todo service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	ListTasks          http.Handler
	CreateTask         http.Handler
	CompleteTask       http.Handler
	RevertTask         http.Handler
	CORS               http.Handler
	GenHTTPOpenapiJSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the todo service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *todo.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemGenHTTPOpenapiJSON http.FileSystem,
) *Server {
	if fileSystemGenHTTPOpenapiJSON == nil {
		fileSystemGenHTTPOpenapiJSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"ListTasks", "GET", "/todo/list"},
			{"CreateTask", "POST", "/todo/task"},
			{"CompleteTask", "POST", "/todo/task/{id}/complete"},
			{"RevertTask", "POST", "/todo/task/{id}/revert"},
			{"CORS", "OPTIONS", "/todo/list"},
			{"CORS", "OPTIONS", "/todo/task"},
			{"CORS", "OPTIONS", "/todo/task/{id}/complete"},
			{"CORS", "OPTIONS", "/todo/task/{id}/revert"},
			{"CORS", "OPTIONS", "/openapi.json"},
			{"./gen/http/openapi.json", "GET", "/openapi.json"},
		},
		ListTasks:          NewListTasksHandler(e.ListTasks, mux, decoder, encoder, errhandler, formatter),
		CreateTask:         NewCreateTaskHandler(e.CreateTask, mux, decoder, encoder, errhandler, formatter),
		CompleteTask:       NewCompleteTaskHandler(e.CompleteTask, mux, decoder, encoder, errhandler, formatter),
		RevertTask:         NewRevertTaskHandler(e.RevertTask, mux, decoder, encoder, errhandler, formatter),
		CORS:               NewCORSHandler(),
		GenHTTPOpenapiJSON: http.FileServer(fileSystemGenHTTPOpenapiJSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "todo" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListTasks = m(s.ListTasks)
	s.CreateTask = m(s.CreateTask)
	s.CompleteTask = m(s.CompleteTask)
	s.RevertTask = m(s.RevertTask)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return todo.MethodNames[:] }

// Mount configures the mux to serve the todo endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListTasksHandler(mux, h.ListTasks)
	MountCreateTaskHandler(mux, h.CreateTask)
	MountCompleteTaskHandler(mux, h.CompleteTask)
	MountRevertTaskHandler(mux, h.RevertTask)
	MountCORSHandler(mux, h.CORS)
	MountGenHTTPOpenapiJSON(mux, goahttp.Replace("", "/./gen/http/openapi.json", h.GenHTTPOpenapiJSON))
}

// Mount configures the mux to serve the todo endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountListTasksHandler configures the mux to serve the "todo" service
// "listTasks" endpoint.
func MountListTasksHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleTodoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/todo/list", f)
}

// NewListTasksHandler creates a HTTP handler which loads the HTTP request and
// calls the "todo" service "listTasks" endpoint.
func NewListTasksHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeListTasksResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listTasks")
		ctx = context.WithValue(ctx, goa.ServiceKey, "todo")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateTaskHandler configures the mux to serve the "todo" service
// "createTask" endpoint.
func MountCreateTaskHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleTodoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/todo/task", f)
}

// NewCreateTaskHandler creates a HTTP handler which loads the HTTP request and
// calls the "todo" service "createTask" endpoint.
func NewCreateTaskHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateTaskRequest(mux, decoder)
		encodeResponse = EncodeCreateTaskResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "createTask")
		ctx = context.WithValue(ctx, goa.ServiceKey, "todo")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCompleteTaskHandler configures the mux to serve the "todo" service
// "completeTask" endpoint.
func MountCompleteTaskHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleTodoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/todo/task/{id}/complete", f)
}

// NewCompleteTaskHandler creates a HTTP handler which loads the HTTP request
// and calls the "todo" service "completeTask" endpoint.
func NewCompleteTaskHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCompleteTaskRequest(mux, decoder)
		encodeResponse = EncodeCompleteTaskResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "completeTask")
		ctx = context.WithValue(ctx, goa.ServiceKey, "todo")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRevertTaskHandler configures the mux to serve the "todo" service
// "revertTask" endpoint.
func MountRevertTaskHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleTodoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/todo/task/{id}/revert", f)
}

// NewRevertTaskHandler creates a HTTP handler which loads the HTTP request and
// calls the "todo" service "revertTask" endpoint.
func NewRevertTaskHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRevertTaskRequest(mux, decoder)
		encodeResponse = EncodeRevertTaskResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "revertTask")
		ctx = context.WithValue(ctx, goa.ServiceKey, "todo")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", HandleTodoOrigin(h).ServeHTTP)
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service todo.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleTodoOrigin(h)
	mux.Handle("OPTIONS", "/todo/list", h.ServeHTTP)
	mux.Handle("OPTIONS", "/todo/task", h.ServeHTTP)
	mux.Handle("OPTIONS", "/todo/task/{id}/complete", h.ServeHTTP)
	mux.Handle("OPTIONS", "/todo/task/{id}/revert", h.ServeHTTP)
	mux.Handle("OPTIONS", "/openapi.json", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleTodoOrigin applies the CORS response headers corresponding to the
// origin for the service todo.
func HandleTodoOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}