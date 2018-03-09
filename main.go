package main

import (
	"fmt"
	"github.com/haraldh/go-varlink-example/orgexamplemore"
	"github.com/varlink/go/varlink"
	"os"
	"time"
)

type more struct {
	// orgexamplemore.VarlinkInterface is added to return
	// MethodNotImplemented for not yet implemented methods
	orgexamplemore.VarlinkInterface
	mycounter int64
	moredata  string
}

func (m *more) Ping(call orgexamplemore.VarlinkCall, ping string) error {
	m.mycounter++
	pong := fmt.Sprintf("%d: %s", m.mycounter, ping)
	return call.ReplyPing(pong)
}

func (m *more) StopServing(call orgexamplemore.VarlinkCall) error {
	service.Stop()
	return call.ReplyStopServing()
}

func (m *more) TestMore(call orgexamplemore.VarlinkCall, n int64) error {
	if !call.WantsMore() {
		return call.ReplyInvalidParameter("more")
	}

	if n > 10 {
		return call.ReplyTestMoreError("n is too big")
	}

	var err error

	call.Continues = true

	err = call.Reply(&struct {
		State interface{} `json:"state"`
	}{State: struct {
		Start bool `json:"start"`
	}{Start: true}})

	if err != nil {
		return err
	}

	for i := int64(0); i < n; i++ {
		err = call.Reply(&struct {
			State interface{} `json:"state"`
		}{State: struct {
			Progress int64 `json:"progress"`
		}{Progress: int64(i * 100 / n)}})

		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	err = call.Reply(&struct {
		State interface{} `json:"state"`
	}{State: struct {
		Progress int64 `json:"progress"`
	}{Progress: int64(100)}})

	if err != nil {
		return err
	}

	call.Continues = false

	return call.Reply(&struct {
		State interface{} `json:"state"`
	}{State: struct {
		Start bool `json:"end"`
	}{Start: true}})
}

func help(name string) {
	fmt.Printf("Usage: %s <varlink address URL>\n", name)
	os.Exit(1)
}

// global only for the method StopServing
var service *varlink.Service

func main() {
	m := more{mycounter: 1, moredata: "test"}

	service = varlink.NewService(
		"Varlink",
		"Example",
		"1",
		"https://github.com/haraldh/go-varlink-example",
	)

	service.RegisterInterface(orgexamplemore.VarlinkNew(&m))

	m.mycounter = 2

	if len(os.Args) < 2 {
		help(os.Args[0])
	}

	err := service.Run(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
}
