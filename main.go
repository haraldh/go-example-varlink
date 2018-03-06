package main

import (
	"fmt"
	"github.com/haraldh/go-varlink-example/orgexamplemore"
	"github.com/varlink/go/varlink"
	"os"
	"time"
)

type more struct {
	varlink.Interface
	mycounter int64
	moredata  string
}

func (m *more) TestMore(call varlink.Call) error {
	if !call.WantsMore() {
		return call.ReplyError("org.varlink.service.InvalidParameter",
			varlink.InvalidParameter_Error{Parameter: "more"})
	}

	var in orgexamplemore.TestMore_In
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

func (m *more) StopServing(call varlink.Call) error {
	service.Stop()
	return call.Reply(&varlink.ServiceOut{})
}

func (m *more) Ping(call varlink.Call) error {
	var in orgexamplemore.Ping_In

	err := call.GetParameters(&in)
	if err != nil {
		return call.ReplyError("org.varlink.service.InvalidParameter", varlink.InvalidParameter_Error{Parameter: "parameters"})
	}

	return call.Reply(&varlink.ServiceOut{
		Parameters: orgexamplemore.Ping_Out{
			in.Ping,
		},
	})
}

func help(name string) {
	fmt.Printf("Usage: %s <varlink address URL>\n", name)
	os.Exit(1)
}

// global only for the method StopServing
var service varlink.Service

func main() {
	m := more{Interface: orgexamplemore.New()}

	service = varlink.NewService(
		"Varlink",
		"Example",
		"1",
		"https://github.com/haraldh/go-varlink-example",
	)
	service.RegisterInterface(&m)

	if len(os.Args) < 2 {
		help(os.Args[0])
	}

	err := service.Run(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
}
