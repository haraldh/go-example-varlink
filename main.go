package main

//go:generate go generate github.com/varlink/go-varlink
//go:generate $GOPATH/bin/varlink-go-generator main ./org.varlink.example.more.varlink

import (
	"github.com/varlink/go-varlink"
	"os"
	"fmt"
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

	service.Run(os.Args[1])
}
