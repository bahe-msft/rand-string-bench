package main

import (
	"fmt"
	"testing"
)

type StringGenerator interface {
	Generate(length int) (string, error)
}

func benchGenerator(b *testing.B, name string, g StringGenerator, length int) {
	b.Run(name, func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, err := g.Generate(length)
			if err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_Generator(b *testing.B) {
	for _, length := range []int{8, 16, 40, 50, 72} {
		b.Run(fmt.Sprintf("length=%d", length), func(b *testing.B) {
			benchGenerator(b, "fixed", NewFixedStringGenerator(), length)
			benchGenerator(b, "math", NewMathRandStringGenerator(), length)
			benchGenerator(b, "crypto", NewCryptoStringGenerator(), length)
		})
	}
}
