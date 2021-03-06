// Code generated by goa v3.7.0, DO NOT EDIT.
//
// todo HTTP client CLI support package
//
// Command:
// $ goa gen github.com/takahashis-shun/todo-goa/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	todo "github.com/takahashis-shun/todo-goa/gen/todo"
)

// BuildHelloPayload builds the payload for the todo hello endpoint from CLI
// flags.
func BuildHelloPayload(todoHelloName string) (*todo.HelloPayload, error) {
	var name string
	{
		name = todoHelloName
	}
	v := &todo.HelloPayload{}
	v.Name = name

	return v, nil
}

// BuildShowPayload builds the payload for the todo show endpoint from CLI
// flags.
func BuildShowPayload(todoShowID string) (*todo.ShowPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(todoShowID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &todo.ShowPayload{}
	v.ID = id

	return v, nil
}

// BuildCreatePayload builds the payload for the todo create endpoint from CLI
// flags.
func BuildCreatePayload(todoCreateBody string) (*todo.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(todoCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"title\": \"Corrupti odit enim nisi itaque eveniet.\"\n   }'")
		}
	}
	v := &todo.CreatePayload{
		Title: body.Title,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the todo update endpoint from CLI
// flags.
func BuildUpdatePayload(todoUpdateBody string, todoUpdateID string) (*todo.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(todoUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"is_done\": false\n   }'")
		}
	}
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(todoUpdateID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &todo.UpdatePayload{
		IsDone: body.IsDone,
	}
	v.ID = id

	return v, nil
}

// BuildDeletePayload builds the payload for the todo delete endpoint from CLI
// flags.
func BuildDeletePayload(todoDeleteID string) (*todo.DeletePayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(todoDeleteID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &todo.DeletePayload{}
	v.ID = id

	return v, nil
}
