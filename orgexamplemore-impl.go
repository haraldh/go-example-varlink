package main

import (
	"encoding/json"

	"github.com/varlink/go-varlink"
)

type Service struct {
	myservice varlink.InterfaceImpl
}

func TestMore(iface *varlink.Interface, call varlink.ServerCall, out *varlink.Writer) error {
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

func StopServing(iface *varlink.Interface, call varlink.ServerCall, out *varlink.Writer) error {
	var in StopServing_CallParameters
	err := json.Unmarshal(*call.Parameters, &in)
	if err != nil {
		return varlink.InvalidParameter("parameters", out)
	}

    // FIXME: Fill me in

	retval := StopServing_ReplyParameters{
		// FIXME: Fill me in
	}

	return out.Reply(varlink.ServerReply{
		Parameters: retval,
	})
}

func Ping(iface *varlink.Interface, call varlink.ServerCall, out *varlink.Writer) error {
	var in Ping_CallParameters
	err := json.Unmarshal(*call.Parameters, &in)
	if err != nil {
		return varlink.InvalidParameter("parameters", out)
	}

    // FIXME: Fill me in

	retval := Ping_ReplyParameters{
		// FIXME: Fill me in
	}

	return out.Reply(varlink.ServerReply{
		Parameters: retval,
	})
}
