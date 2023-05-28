package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("todo", func() {
	Title("todo app")
	Description("application for management tasks")

	Server("todo", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var Task = Type("Task", func() {
	Field(10, "id", String, "id of task")
	Field(11, "name", String, "name of task")
	Field(12, "completed", Boolean, "status of task")
})

var _ = Service("todo", func() {
	Description("todo service for management tasks")

	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("Content-Type")
		cors.Methods("GET", "POST", "OPTIONS")
	})

	Method("listTasks", func() {
		Result(ArrayOf(Task))

		HTTP(func() {
			GET("/todo/list")
		})
	})

	Method("createTask", func() {
		Payload(func() {
			Field(1, "name", String, "name of task")
		})

		Result(Task)

		HTTP(func() {
			POST("/todo/task")
		})
		GRPC(func() {
		})
	})

	Method("completeTask", func() {
		Payload(func() {
			Field(2, "id", String, "id of task")
		})

		Result(Task)

		HTTP(func() {
			POST("/todo/task/{id}/complete")
		})
		GRPC(func() {
		})
	})

	Method("revertTask", func() {
		Payload(func() {
			Field(3, "id", String, "id of task")
		})

		Result(Task)

		HTTP(func() {
			POST("/todo/task/{id}/revert")
		})
		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
