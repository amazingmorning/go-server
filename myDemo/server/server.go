package main

import (
	"go-server/gzet"
)

func main() {
	s := gzet.NewServer("first-server")
	s.Serve()
}
