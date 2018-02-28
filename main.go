package main

//go:generate go generate ../go-varlink

import (
	"../go-varlink"
	"os"
)

func main() {
	ifaces := []varlink.VarlinkInterface{
		// list own interface here
	}
	service := varlink.NewService(
		"Atomic",
		"podman",
		"0.5",
		"https://github.com/projectatomic/libpod",
		ifaces,
	)

	service.Run(os.Args[1])
}
