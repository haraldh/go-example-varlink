package main

//go:generate go generate ../go-varlink

import (
	"../go-varlink"
	"os"
)

func main() {
	service := varlink.NewService(
		"Atomic",
		"podman",
		"0.5",
		"https://github.com/projectatomic/libpod",
		[]varlink.VarlinkInterface{},
	)

	// Register a new interface
	// myiface = StructWithVarlinkInterface {...}
	//service.RegisterInterface(VarlinkInterface(myiface))

	service.Run(os.Args[1])
}
