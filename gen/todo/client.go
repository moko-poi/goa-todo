// Code generated by goa v3.7.0, DO NOT EDIT.
//
// todo client
//
// Command:
// $ goa gen github.com/takahashis-shun/todo-goa/design

package todo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "todo" service client.
type Client struct {
	HelloEndpoint  goa.Endpoint
	ShowEndpoint   goa.Endpoint
	CreateEndpoint goa.Endpoint
}

// NewClient initializes a "todo" service client given the endpoints.
func NewClient(hello, show, create goa.Endpoint) *Client {
	return &Client{
		HelloEndpoint:  hello,
		ShowEndpoint:   show,
		CreateEndpoint: create,
	}
}

// Hello calls the "hello" endpoint of the "todo" service.
func (c *Client) Hello(ctx context.Context, p *HelloPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.HelloEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Show calls the "show" endpoint of the "todo" service.
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *Todo, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Todo), nil
}

// Create calls the "create" endpoint of the "todo" service.
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res string, err error) {
	var ires interface{}
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}