package main

import (
	"fmt"
	"testing"
)

var states = []struct {
	input string
}{
	{input: "solid"},
	{input: "liquid"},
	{input: "gas"},
	{input: "plasma"},
}

func BenchmarkIfElseFunc(b *testing.B) {
	for _, v := range states {
		b.Run(fmt.Sprintf("input_size_%s", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ifElseFn(v.input)
			}
		})
	}
}

func BenchmarkSwitchFunc(b *testing.B) {
	for _, v := range states {
		b.Run(fmt.Sprintf("input_size_%s", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				switchFn(v.input)
			}
		})
	}
}
