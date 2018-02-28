package main

import (
	"github.com/varlink/go-varlink"
)

type Ping_CallParameters struct {
	Ping string `json:"ping"`
}

type Ping_ReplyParameters struct {
	Pong string `json:"pong"`
}

type State struct {
	Start    bool  `json:"start"`
	Progress int64 `json:"progress"`
	End      int64 `json:"end"`
}

type TestMore_CallParameters struct {
	N int64 `json:"n"`
}

type TestMore_ReplyParameters struct {
	State State `json:"state"`
}

type StopServing_CallParameters struct {
}

type StopServing_ReplyParameters struct {
}

func (this Service) Get() *varlink.InterfaceImpl {
	return &this.myservice
}

func NewService() Service {
	r := Service{
		myservice: varlink.InterfaceImpl{
			Name: "org.example.more",
			Description: `# Example Varlink service
interface org.example.more

# Enum, returning either start, progress or end
# progress: [0-100]
type State (
     start: bool,
     progress: int,
     end: bool
)

# Returns the same string
method Ping(ping : string) -> (pong: string)

# Dummy progress method
# n: number of progress steps
method TestMore(n : int) -> (state: State)

# Stop serving
method StopServing() -> ()

# Something failed
error ActionFailed (reason: string)
`,
			Methods: map[string]varlink.Method{
				"Ping":        Ping,
				"TestMore":    TestMore,
				"StopServing": StopServing,
			},
		},
	}
	return r
}
