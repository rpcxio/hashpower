package main

import (
	"flag"

	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", ":8972", "listened address")

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Hash", new(HashPower), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
