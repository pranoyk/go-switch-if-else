## GO-SWITCH-IF-ELSE

The aim of this project is to figure out the difference in performance in between SWITCH statements and IF-ELSE statements.

---

On adding benchmark tests for `switchFn` and `ifElseFn` functions I found the below results. There seems to be no significant difference for simple if else and switch commands.

Command run - 
```
go test -bench=.
```
Output - 
```
goos: darwin
goarch: arm64
pkg: github.com/pranoyk/go-switch-if-else
cpu: Apple M3 Pro
BenchmarkIfElseFunc/input_size_solid-11         	105697726	        11.16 ns/op
BenchmarkIfElseFunc/input_size_liquid-11        	100000000	        11.32 ns/op
BenchmarkIfElseFunc/input_size_gas-11           	100000000	        11.43 ns/op
BenchmarkIfElseFunc/input_size_plasma-11        	100000000	        11.56 ns/op
BenchmarkSwitchFunc/input_size_solid-11         	100000000	        11.22 ns/op
BenchmarkSwitchFunc/input_size_liquid-11        	100000000	        11.45 ns/op
BenchmarkSwitchFunc/input_size_gas-11           	96215522	        11.03 ns/op
BenchmarkSwitchFunc/input_size_plasma-11        	96402320	        11.51 ns/op
PASS
ok  	github.com/pranoyk/go-switch-if-else	9.932s
```

go test with memory allocation - `go test -bench=. -benchmem`
Output -
```
goos: darwin
goarch: arm64
pkg: github.com/pranoyk/go-switch-if-else
cpu: Apple M3 Pro
BenchmarkIfElseFunc/input_size_solid-11         	107347539	        11.28 ns/op	       8 B/op	       1 allocs/op
BenchmarkIfElseFunc/input_size_liquid-11        	92172976	        11.23 ns/op	       8 B/op	       1 allocs/op
BenchmarkIfElseFunc/input_size_gas-11           	98018858	        11.57 ns/op	       8 B/op	       1 allocs/op
BenchmarkIfElseFunc/input_size_plasma-11        	97315376	        11.69 ns/op	       8 B/op	       1 allocs/op
BenchmarkSwitchFunc/input_size_solid-11         	100000000	        11.25 ns/op	       8 B/op	       1 allocs/op
BenchmarkSwitchFunc/input_size_liquid-11        	99550635	        11.53 ns/op	       8 B/op	       1 allocs/op
BenchmarkSwitchFunc/input_size_gas-11           	100000000	        11.11 ns/op	       8 B/op	       1 allocs/op
BenchmarkSwitchFunc/input_size_plasma-11        	90286374	        11.68 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/pranoyk/go-switch-if-else	9.812s
```