package main

//go:generate go generate github.com/varlink/go-varlink

import (
	"github.com/varlink/go-varlink"
	"os"
)

func main() {
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

	service.Run(os.Args[1])
}
