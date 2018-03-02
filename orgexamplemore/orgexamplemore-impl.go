package orgexamplemore

import (
	"encoding/json"
	"github.com/varlink/go-varlink"
	"reflect"
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

func (this *Interface) HandleOld(method string, call varlink.ServerCall, out *varlink.Writer) error {
	switch method {
	case "Ping":
		return this.Ping(call, out)
	case "StopServing":
		return this.StopServing(call, out)
	case "TestMore":
		return this.TestMore(call, out)
	}
	return varlink.MethodNotFound(method, out)
}

func (this *Interface) Handle(method string, call varlink.ServerCall, out *varlink.Writer) error {
	// MethodByName() returns 'zero Kind' for unknown methods
	v := reflect.ValueOf(this).MethodByName(method)
	if v.Kind() != reflect.Func {
		return varlink.MethodNotFound(method, out)
	}

	args := []reflect.Value{
		reflect.ValueOf(call),
		reflect.ValueOf(out),
	}
	ret := v.Call(args)

	if ret[0].Interface() == nil {
		return nil
	}

	return ret[0].Interface().(error)
}

func NewInterface() Interface {
	r := Interface{
		InterfaceDefinition: varlink.InterfaceDefinition{
			Name:        "org.example.more",
			Description: InterfaceDescription,
			Methods:     []string{"Ping", "StopServing", "TestMore"},
		},
	}
	return r
}
