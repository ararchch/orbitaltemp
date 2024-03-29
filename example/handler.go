package main

import (
	"context"
	api "example/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// returns 'Hello followed by the incoming req message
	return &api.Response{Message: "Hello, " + req.Message}, nil
}
