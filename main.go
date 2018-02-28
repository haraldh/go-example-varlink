package main

//go:generate go generate github.com/varlink/go-varlink

import (
	"github.com/varlink/go-varlink"
	"os"
)

func main() {
	var service = varlink.NewService(
		"Atomic",
		"podman",
		"0.5",
		"https://github.com/projectatomic/libpod",
	)

	// Register a new interface
	// myiface = StructWithVarlinkInterface {...}
	//service.RegisterInterface(VarlinkInterface(myiface))

	service.Run(os.Args[1])
}
