package main

import (
	"fmt"
	"github.com/haraldh/go-varlink-example/orgexamplemore"
	"github.com/varlink/go/varlink"
	"os"
	"sync"
	"time"
)

type more struct {
	// orgexamplemore.VarlinkInterface is added to return
	// MethodNotImplemented for not yet implemented methods
	// If it is omitted, the compiler will check, if the implementation
	// is complete.
	orgexamplemore.VarlinkInterface
	sync.RWMutex
	mycounter int64
	moredata  string
}

// Ping returns the given ping string and adds an increasing counter
// from the global more struct, which is guarded with a sync.RWMutex
func (m *more) Ping(call orgexamplemore.VarlinkCall, ping string) error {
	m.Lock()
	m.mycounter++
	m.Unlock()

	m.RLock()
	pong := fmt.Sprintf("%d: %s", m.mycounter, ping)
	m.RUnlock()

	return call.ReplyPing(pong)
}

func (m *more) StopServing(call orgexamplemore.VarlinkCall) error {
	service.Shutdown()
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

	bstart := true
	start := orgexamplemore.State{Start: &bstart}

	err = call.ReplyTestMore(start)

	if err != nil {
		return err
	}

	for i := int64(0); i < n; i++ {
		p := int64(i * 100 / n)
		progress := orgexamplemore.State{Progress: &p}

		err = call.ReplyTestMore(progress)

		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	p := int64(100)
	progress := orgexamplemore.State{Progress: &p}

	err = call.ReplyTestMore(progress)

	if err != nil {
		return err
	}

	call.Continues = false

	bend := true
	end := orgexamplemore.State{End: &bend}

	return call.ReplyTestMore(end)
}

func help(name string) {
	fmt.Printf("Usage: %s <varlink address URL>\n", name)
	os.Exit(1)
}

// global only for the method StopServing
var service *varlink.Service

func main() {
	m := more{mycounter: 1, moredata: "test"}

	service, _ = varlink.NewService(
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

	err := service.Listen(os.Args[1], time.Duration(10)*time.Second)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
