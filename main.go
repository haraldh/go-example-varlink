package main

import (
	"fmt"
	"github.com/haraldh/go-varlink-example/orgexamplemore"
	"github.com/varlink/go-varlink"
	"os"
)

func help(name string) {
	fmt.Printf("Usage: %s <varlink address URL>\n", name)
	os.Exit(1)
}

func main() {
	orgexamplemoreiface := orgexamplemore.NewInterface()
	service := varlink.NewService(
		"Varlink",
		"Example",
		"1",
		"https://github.com/haraldh/go-varlink-example",
		[]varlink.Interface{
			&orgexamplemoreiface,
		},
	)

	if len(os.Args) < 2 {
		help(os.Args[0])
	}

	// fill in extra data, for the StopServing() method
	orgexamplemoreiface.Server = &service

	err := service.Run(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
}
