## GO-SWITCH-IF-ELSE

The aim of this project is to figure out the difference in performance in between SWITCH statements and IF-ELSE statements.

---

On adding benchmark tests for `switchFn` and `ifElseFn` functions I found the below results. There seems to be no significant difference for simple if else and switch commands.

Command run - 
```
go test -bench=. -count 5
```
Output - 
```
goos: darwin
goarch: arm64
pkg: github.com/pranoyk/go-switch-if-else
cpu: Apple M3 Pro
BenchmarkIfElseFnSolidState-11     	111179576	        10.57 ns/op
BenchmarkIfElseFnSolidState-11     	99706422	        10.73 ns/op
BenchmarkIfElseFnSolidState-11     	100000000	        10.59 ns/op
BenchmarkIfElseFnSolidState-11     	100000000	        10.57 ns/op
BenchmarkIfElseFnSolidState-11     	100000000	        10.61 ns/op
BenchmarkSwitchFnSolidState-11     	100000000	        10.62 ns/op
BenchmarkSwitchFnSolidState-11     	100000000	        10.61 ns/op
BenchmarkSwitchFnSolidState-11     	100000000	        10.60 ns/op
BenchmarkSwitchFnSolidState-11     	100000000	        10.59 ns/op
BenchmarkSwitchFnSolidState-11     	100000000	        10.59 ns/op
BenchmarkIfElseFnLiquidState-11    	100000000	        10.60 ns/op
BenchmarkIfElseFnLiquidState-11    	100000000	        10.55 ns/op
BenchmarkIfElseFnLiquidState-11    	100000000	        10.62 ns/op
BenchmarkIfElseFnLiquidState-11    	100000000	        10.79 ns/op
BenchmarkIfElseFnLiquidState-11    	100000000	        10.74 ns/op
BenchmarkSwitchFnLiquidState-11    	100000000	        10.66 ns/op
BenchmarkSwitchFnLiquidState-11    	100000000	        10.63 ns/op
BenchmarkSwitchFnLiquidState-11    	100000000	        10.65 ns/op
BenchmarkSwitchFnLiquidState-11    	100000000	        10.65 ns/op
BenchmarkSwitchFnLiquidState-11    	100000000	        10.67 ns/op
BenchmarkIfElseFnGasState-11       	100000000	        10.71 ns/op
BenchmarkIfElseFnGasState-11       	100000000	        10.65 ns/op
BenchmarkIfElseFnGasState-11       	100000000	        10.60 ns/op
BenchmarkIfElseFnGasState-11       	100000000	        10.61 ns/op
BenchmarkIfElseFnGasState-11       	100000000	        10.60 ns/op
BenchmarkSwitchFnGasState-11       	100000000	        10.79 ns/op
BenchmarkSwitchFnGasState-11       	100000000	        10.86 ns/op
BenchmarkSwitchFnGasState-11       	98683195	        10.86 ns/op
BenchmarkSwitchFnGasState-11       	100000000	        10.86 ns/op
BenchmarkSwitchFnGasState-11       	100000000	        10.85 ns/op
BenchmarkIfElseFnPlasmaState-11    	100000000	        10.87 ns/op
BenchmarkIfElseFnPlasmaState-11    	100000000	        10.94 ns/op
BenchmarkIfElseFnPlasmaState-11    	100000000	        10.95 ns/op
BenchmarkIfElseFnPlasmaState-11    	100000000	        10.90 ns/op
BenchmarkIfElseFnPlasmaState-11    	100000000	        10.76 ns/op
BenchmarkSwitchFnPlasmaState-11    	100000000	        10.72 ns/op
BenchmarkSwitchFnPlasmaState-11    	100000000	        10.67 ns/op
BenchmarkSwitchFnPlasmaState-11    	100000000	        11.01 ns/op
BenchmarkSwitchFnPlasmaState-11    	100000000	        10.94 ns/op
BenchmarkSwitchFnPlasmaState-11    	100000000	        10.84 ns/op
PASS
ok  	github.com/pranoyk/go-switch-if-else	44.578s
```