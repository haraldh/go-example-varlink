package main

import (
	"fmt"
	"github.com/varlink/go-varlink"
	"os"
	"github.com/haraldh/go-varlink-example/orgexamplemore"
)

func help(name string) {
	fmt.Printf("Usage: %s <varlink address URL>\n", name)
	os.Exit(1)
}

func main() {
	myservice := orgexamplemore.NewService()
	service := varlink.NewService(
		"Varlink",
		"Example",
		"1",
		"https://github.com/haraldh/go-varlink-example",
		[]varlink.Interface{
			&myservice,
		},
	)

	if len(os.Args) < 2 {
		help(os.Args[0])
	}
	myservice.Server = &service

	err := service.Run(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
}
