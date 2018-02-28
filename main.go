package main

//go:generate $GOPATH/bin/varlink-generator main ./org.varlink.example.more.varlink

import (
	"fmt"
	"os"

	"github.com/varlink/go-varlink"
)

func help(name string) {
	fmt.Printf("Usage: %s <varlink address URL>\n", name)
	os.Exit(1)
}

func main() {
	fmt.Println(OrgVarlinkExampleMore)
	ifaces := []varlink.Interface{
		// list own interface here
	}
	service := varlink.NewService(
		"Atomic",
		"podman",
		"0.5",
		"https://github.com/projectatomic/libpod",
		ifaces,
	)

	if len(os.Args) < 2 {
		help(os.Args[0])
	}

	err := service.Run(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
}
