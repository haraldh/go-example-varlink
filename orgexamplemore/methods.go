package orgexamplemore

//go:generate $GOPATH/bin/varlink-generator ./org.example.more.varlink

import (
	"github.com/varlink/go-varlink"
)

type Interface struct {
	varlink.InterfaceDefinition
	Server *varlink.Service
}

func NewInterface() Interface {
	return Interface{InterfaceDefinition: NewInterfaceDefinition()}
}

func (this *Interface) TestMore(ctx varlink.Context) error {
	var in TestMore_In
	err := ctx.Args(&in)
	if err != nil {
		return ctx.Reply(&varlink.ServerOut{
			Error:      "org.varlink.service.InvalidParameter",
			Parameters: varlink.InvalidParameter_Error{Parameter: "parameters"},
		})
	}

	// FIXME: Fill me in
	return ctx.Reply(&varlink.ServerOut{
		Error:      "org.varlink.service.MethodNotImplemented",
		Parameters: varlink.MethodNotImplemented_Error{Method: "TestMore"},
	})

	return ctx.Reply(&varlink.ServerOut{
		Parameters: TestMore_Out{
		// FIXME: Fill me in
		},
	})
}

func (this *Interface) StopServing(ctx varlink.Context) error {
	if this.Server != nil {
		this.Server.Stop()
	}
	return ctx.Reply(&varlink.ServerOut{})
}

func (this *Interface) Ping(ctx varlink.Context) error {
	var in Ping_In

	err := ctx.Args(&in)
	if err != nil {
		return ctx.Reply(&varlink.ServerOut{
			Error:      "org.varlink.service.InvalidParameter",
			Parameters: varlink.InvalidParameter_Error{Parameter: "parameters"},
		})
	}

	return ctx.Reply(&varlink.ServerOut{
		Parameters: Ping_Out{
			in.Ping,
		},
	})
}

func ActionFailed(ctx varlink.Context, reason string) error {
	return ctx.Reply(&varlink.ServerOut{
		Error: "org.example.more.ActionFailed",
		Parameters: ActionFailed_Error{
			Reason: reason,
		},
	})
}
