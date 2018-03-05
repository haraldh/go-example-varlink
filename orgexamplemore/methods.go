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
	var in TestMore_In
	err := call.GetParameters(&in)
	if err != nil {
		return call.ReplyError("org.varlink.service.InvalidParameter", varlink.InvalidParameter_Error{Parameter: "parameters"})
	}

	call.Reply(&varlink.ServiceOut{
		Continues:  true,
		Parameters: TestMore_Out{State: State{Start: true}},
	})

	for i := int64(0); i < in.N; i++ {
		call.Reply(&varlink.ServiceOut{
			Continues:  true,
			Parameters: TestMore_Out{State: State{Progress: int64(i * 100 / in.N)}},
		})
		time.Sleep(time.Second)
	}
	call.Reply(&varlink.ServiceOut{
		Continues:  true,
		Parameters: TestMore_Out{State: State{Progress: int64(100)}},
	})

	return call.Reply(&varlink.ServiceOut{
		Continues:  false,
		Parameters: TestMore_Out{State: State{End: true}},
	})

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
