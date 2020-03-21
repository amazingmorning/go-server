package main

import (
	"go-server/gnet"
)

func main() {
	s := gnet.NewServer("first-server")
	s.Serve()
}
