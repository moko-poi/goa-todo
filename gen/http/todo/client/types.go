// Code generated by goa v3.7.0, DO NOT EDIT.
//
// todo HTTP client types
//
// Command:
// $ goa gen github.com/takahashis-shun/todo-goa/design

package client

import (
	todo "github.com/takahashis-shun/todo-goa/gen/todo"
	todoviews "github.com/takahashis-shun/todo-goa/gen/todo/views"
)

// CreateRequestBody is the type of the "todo" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Title
	Title string `form:"title" json:"title" xml:"title"`
}

// UpdateRequestBody is the type of the "todo" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// IsDone
	IsDone bool `form:"is_done" json:"is_done" xml:"is_done"`
}

// ShowResponseBody is the type of the "todo" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// IsDone
	IsDone *bool `form:"is_done,omitempty" json:"is_done,omitempty" xml:"is_done,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "todo" service.
func NewCreateRequestBody(p *todo.CreatePayload) *CreateRequestBody {
	body := &CreateRequestBody{
		Title: p.Title,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "todo" service.
func NewUpdateRequestBody(p *todo.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		IsDone: p.IsDone,
	}
	return body
}

// NewShowTodoOK builds a "todo" service "show" endpoint result from a HTTP
// "OK" response.
func NewShowTodoOK(body *ShowResponseBody) *todoviews.TodoView {
	v := &todoviews.TodoView{
		ID:     body.ID,
		Title:  body.Title,
		IsDone: body.IsDone,
	}

	return v
}
