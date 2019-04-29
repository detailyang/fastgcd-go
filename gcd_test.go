package fastgcd

import (
	"testing"

	"math/big"
	"math/rand"
)

func TestGCD(t *testing.T) {
	for i := 1; i < 10000000; i++ {
		y := rand.Intn(i) + 1
		x := GCD(uint64(i), uint64(y))
		w := new(big.Int)
		w.GCD(nil, nil, big.NewInt(int64(x)), big.NewInt(int64(y)))
		if w.Int64() != int64(x) {
			t.Fatalf("%d, %d expect %d but got %d", i, y, w.Int64(), x)
		}
	}
}

func TestBinaryGCD(t *testing.T) {
	for i := 1; i < 10000000; i++ {
		y := rand.Intn(i) + 1
		x := BinaryGCD(uint64(i), uint64(y))
		w := new(big.Int)
		w.GCD(nil, nil, big.NewInt(int64(x)), big.NewInt(int64(y)))
		if w.Int64() != int64(x) {
			t.Fatalf("%d, %d expect %d but got %d", i, y, w.Int64(), x)
		}
	}
}
