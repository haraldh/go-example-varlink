package orgexamplemore

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

func NewService() Service {
	r := Service{
		InterfaceImpl: varlink.InterfaceImpl{
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
		},
	}
	return r
}

func (this *Service) Handle(method string, call varlink.ServerCall, out *varlink.Writer) error {
	switch method {
	case "Ping":
		return this.Ping(call, out)
	case "TestMore":
		return this.TestMore(call, out)
	case "StopServing":
		return this.StopServing(call, out)
	}
	return varlink.MethodNotFound(method, out)
}
