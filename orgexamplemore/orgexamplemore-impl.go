package orgexamplemore

import (
	"encoding/json"
	"github.com/varlink/go-varlink"
)

type Interface struct {
	varlink.InterfaceDefinition
	Server *varlink.Service
}

func (this *Interface) TestMore(call varlink.ServerCall, out *varlink.Writer) error {
	var in TestMore_In
	err := json.Unmarshal(*call.Parameters, &in)
	if err != nil {
		return varlink.InvalidParameter("parameters", out)
	}

	// FIXME: Fill me in
	return varlink.MethodNotImplemented("TestMore", out)

	retval := TestMore_Out{
		// FIXME: Fill me in
	}

	return out.Reply(varlink.ServerReply{
		Parameters: retval,
	})
}

func (this *Interface) StopServing(call varlink.ServerCall, out *varlink.Writer) error {
	if this.Server != nil {
		this.Server.Stop()
	}
	return out.Reply(varlink.ServerReply{})
}

func (this *Interface) Ping(call varlink.ServerCall, out *varlink.Writer) error {
	var in Ping_In
	if call.Parameters == nil {
		return varlink.InvalidParameter("parameters", out)
	}

	err := json.Unmarshal(*call.Parameters, &in)
	if err != nil {
		return varlink.InvalidParameter("parameters", out)
	}

	retval := Ping_Out{
		in.Ping,
	}

	return out.Reply(varlink.ServerReply{
		Parameters: retval,
	})
}

func NewInterface() Interface {
	r := Interface{InterfaceDefinition: NewInterfaceDefinition()}
	r.Methods = []string{"Ping", "StopServing", "TestMore"}
	return r
}
