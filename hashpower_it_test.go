// +build it
// go test -tags it -timeout 1m -run IT .

package main

import (
	"context"
	"testing"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/stretchr/testify/assert"
)

func TestHashPower_Calc_IT(t *testing.T) {
	s := server.NewServer()
	defer s.Close()

	s.RegisterName("Hash", new(HashPower), "")
	go s.Serve("tcp", "127.0.0.1:0")
	time.Sleep(time.Second)

	addr := s.Address().String()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	opt := client.DefaultOption

	xclient := client.NewXClient("Hash", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	var data string
	err := xclient.Call(ctx, "Calc", 10, &data)
	assert.NoError(t, err)
	assert.Equal(t, "0000000000", data[:10])
	cancel()
}
