package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPower_Calc(t *testing.T) {
	hp := &HashPower{}

	var reply string
	err := hp.Calc(context.Background(), 10, &reply)
	assert.NoError(t, err)
	assert.Greater(t, len(reply), 10)
	assert.Equal(t, "0000000000", reply[:10])
}

func BenchmarkHashPower_Calc(b *testing.B) {
	hp := &HashPower{}
	var reply string

	for i := 0; i < b.N; i++ {
		hp.Calc(context.Background(), 10, &reply)
	}
}
