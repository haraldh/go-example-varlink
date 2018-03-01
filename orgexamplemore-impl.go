package main

import (
	"encoding/json"
	"github.com/varlink/go-varlink"
)

type Service struct {
	varlink.InterfaceImpl
	server *varlink.Service
}

func (this *Service) Handle(method string, call varlink.ServerCall, out *varlink.Writer) error {
	switch method {
	case "Ping":
		return this.Ping(call, out)
	case "TestMore":
		return this.TestMore(call, out)
	case "StopServing":
		return this.StopServing(call, out)
	}
	return varlink.MethodNotFound(method, out)
}

func (this *Service) TestMore(call varlink.ServerCall, out *varlink.Writer) error {
	var in TestMore_CallParameters
	err := json.Unmarshal(*call.Parameters, &in)
	if err != nil {
		return varlink.InvalidParameter("parameters", out)
	}

	// FIXME: Fill me in

	retval := TestMore_ReplyParameters{
		// FIXME: Fill me in
	}

	return out.Reply(varlink.ServerReply{
		Parameters: retval,
	})
}

func (this *Service) StopServing(call varlink.ServerCall, out *varlink.Writer) error {
	this.server.Stop()
	return out.Reply(varlink.ServerReply{})
}

func (this *Service) Ping(call varlink.ServerCall, out *varlink.Writer) error {
	var in Ping_CallParameters
	if call.Parameters == nil {
		return varlink.InvalidParameter("parameters", out)
	}

	err := json.Unmarshal(*call.Parameters, &in)
	if err != nil {
		return varlink.InvalidParameter("parameters", out)
	}

	retval := Ping_ReplyParameters{
		in.Ping,
	}

	return out.Reply(varlink.ServerReply{
		Parameters: retval,
	})
}
