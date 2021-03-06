package fastgcd

import (
	"math/big"
	"testing"
)

func BenchmarkVanillaGCD(b *testing.B) {
	z := big.NewInt(0)
	x := big.NewInt(0)
	y := big.NewInt(0)
	for i := uint64(0); i < uint64(b.N); i++ {
		x.SetInt64(int64(i))
		y.SetInt64(int64(i + uint64(b.N)))
		z.GCD(nil, nil, x, y)
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		GCD(i, i+uint64(b.N))
	}
}

func BenchmarkBinaryGCD(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		BinaryGCD(i, i+uint64(b.N))
	}
}
