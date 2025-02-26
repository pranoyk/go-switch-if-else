package main

import (
	"testing"
)


func BenchmarkIfElseFnSolidState(b *testing.B) {
	var state = "solid"
	for i := 0; i < b.N; i++ {
		ifElseFn(state)
	}
}

func BenchmarkSwitchFnSolidState(b *testing.B) {
	var state = "solid"
	for i := 0; i < b.N; i++ {
		switchFn(state)
	}
}

func BenchmarkIfElseFnLiquidState(b *testing.B) {
	var state = "liquid"
	for i := 0; i < b.N; i++ {
		ifElseFn(state)
	}
}

func BenchmarkSwitchFnLiquidState(b *testing.B) {
	var state = "liquid"
	for i := 0; i < b.N; i++ {
		switchFn(state)
	}
}

func BenchmarkIfElseFnGasState(b *testing.B) {
	var state = "gas"
	for i := 0; i < b.N; i++ {
		ifElseFn(state)
	}
}

func BenchmarkSwitchFnGasState(b *testing.B) {
	var state = "gas"
	for i := 0; i < b.N; i++ {
		switchFn(state)
	}
}

func BenchmarkIfElseFnPlasmaState(b *testing.B) {
	var state = "plasma"
	for i := 0; i < b.N; i++ {
		ifElseFn(state)
	}
}

func BenchmarkSwitchFnPlasmaState(b *testing.B) {
	var state = "plasma"
	for i := 0; i < b.N; i++ {
		switchFn(state)
	}
}