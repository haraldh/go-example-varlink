package orgexamplemore

//go:generate $GOPATH/bin/varlink-generator ./org.example.more.varlink

import (
	"github.com/varlink/go/varlink"
	"time"
)

type Interface struct {
	varlink.InterfaceDefinition
	Service *varlink.Service
}

func NewInterface() Interface {
	return Interface{InterfaceDefinition: NewInterfaceDefinition()}
}

func (intf *Interface) TestMore(call varlink.Call) error {
	if !call.WantsMore() {
		return call.ReplyError("org.varlink.service.InvalidParameter",
			varlink.InvalidParameter_Error{Parameter: "more"})
	}

	var in TestMore_In
	if err := call.GetParameters(&in); err != nil {
		return call.ReplyError("org.varlink.service.InvalidParameter",
			varlink.InvalidParameter_Error{Parameter: "parameters"})
	}

	if err := call.Reply(&varlink.ServiceOut{Continues: true,
		Parameters: struct {
			State interface{} `json:"state"`
		}{State: struct {
			Start bool `json:"start"`
		}{Start: true}}}); err != nil {
		return err
	}

	for i := int64(0); i < in.N; i++ {
		if err := call.Reply(&varlink.ServiceOut{
			Continues: true,
			Parameters: struct {
				State interface{} `json:"state"`
			}{State: struct {
				Progress int64 `json:"progress"`
			}{Progress: int64(i * 100 / in.N)}}}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	if err := call.Reply(&varlink.ServiceOut{
		Continues: true,
		Parameters: struct {
			State interface{} `json:"state"`
		}{State: struct {
			Progress int64 `json:"progress"`
		}{Progress: int64(100)}}}); err != nil {
		return err
	}

	return call.Reply(&varlink.ServiceOut{
		Continues: false,
		Parameters: struct {
			State interface{} `json:"state"`
		}{State: struct {
			Start bool `json:"end"`
		}{Start: true}}})
}

func (intf *Interface) StopServing(call varlink.Call) error {
	if intf.Service != nil {
		intf.Service.Stop()
	}
	return call.Reply(&varlink.ServiceOut{})
}

func (intf *Interface) Ping(call varlink.Call) error {
	var in Ping_In

	err := call.GetParameters(&in)
	if err != nil {
		return call.ReplyError("org.varlink.service.InvalidParameter", varlink.InvalidParameter_Error{Parameter: "parameters"})
	}

	return call.Reply(&varlink.ServiceOut{
		Parameters: Ping_Out{
			in.Ping,
		},
	})
}
