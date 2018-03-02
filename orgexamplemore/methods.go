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
	err := ctx.Arguments(&in)
	if err != nil {
		return varlink.InvalidParameter(ctx, "parameters")
	}

	// FIXME: Fill me in
	return varlink.MethodNotImplemented(ctx, "TestMore")

	return ctx.Reply(&varlink.ServerReply{
		Parameters: TestMore_Out{
			// FIXME: Fill me in
		},
	})
}

func (this *Interface) StopServing(ctx varlink.Context) error {
	if this.Server != nil {
		this.Server.Stop()
	}
	return ctx.Reply(&varlink.ServerReply{})
}

func (this *Interface) Ping(ctx varlink.Context) error {
	var in Ping_In

	err := ctx.Arguments(&in)
	if err != nil {
		return varlink.InvalidParameter(ctx, "parameters")
	}

	return ctx.Reply(&varlink.ServerReply{
		Parameters: Ping_Out{
			in.Ping,
		},
	})
}
