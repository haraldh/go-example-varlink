package main

//go:generate $GOPATH/bin/varlink-generator main ./org.example.more.varlink

import (
	"fmt"
	"github.com/varlink/go-varlink"
	"os"
)

func help(name string) {
	fmt.Printf("Usage: %s <varlink address URL>\n", name)
	os.Exit(1)
}

func main() {
	myservice := NewService()
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
	myservice.server = &service

	err := service.Run(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
}
