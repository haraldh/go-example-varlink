package orgexamplemore

import (
	"encoding/json"
	"github.com/varlink/go-varlink"
)

type Service struct {
	varlink.InterfaceImpl
	Server *varlink.Service
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
	if this.Server != nil {
		this.Server.Stop()
	}
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
