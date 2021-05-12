package main

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"strconv"
	"time"
)

// HashPower is a hash service like bitcoin mining.
type HashPower struct{}

func (s *HashPower) Calc(ctx context.Context, target int, value *string) error {
	if target > 10 {
		return errors.New("current machine has no more power")
	}
	*value = pow(target)

	return nil
}

func pow(tb int) string {
	target := big.NewInt(1)
	// 难度：当hash值的前targetBits位均为0即满足要求,
	target.Lsh(target, uint(256-tb))

	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for {
		data := fmt.Sprintf("%d", time.Now().UnixNano()) + strconv.Itoa(nonce)
		hash = sha256.Sum256([]byte(data))
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			break
		} else {
			nonce++
		}

		if nonce%100 == 0 {
			runtime.Gosched()
		}
	}

	return fmt.Sprintf("%0256b", &hashInt)
}
