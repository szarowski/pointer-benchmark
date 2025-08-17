# Pointer Benchmark Test

This benchmark shows the difference in execution of simple update functions with arguments passed by value 
or passed by reference in Golang.

## Run

```
go test -bench=BenchmarkArrayStructs
```

## Results

Results of running the benchmark on MacBook Pro M2 with Apple Silicon architecture and 32 GB memory:

```
goos: darwin
goarch: arm64
pkg: reference-benchmark
cpu: Apple M2 Max
BenchmarkArrayStructs/Value_Size_1-12           82960102                14.17 ns/op
BenchmarkArrayStructs/Pointer_Size_1-12         83931249                14.42 ns/op
BenchmarkArrayStructs/Value_Size_10-12          78162956                15.81 ns/op
BenchmarkArrayStructs/Pointer_Size_10-12        80116617                15.20 ns/op
BenchmarkArrayStructs/Value_Size_100-12         44019870                28.37 ns/op
BenchmarkArrayStructs/Pointer_Size_100-12       78062955                15.32 ns/op
BenchmarkArrayStructs/Value_Size_1000-12         6975787               171.6 ns/op
BenchmarkArrayStructs/Pointer_Size_1000-12      80187997                15.20 ns/op
BenchmarkArrayStructs/Value_Size_10000-12         533920              2461 ns/op
BenchmarkArrayStructs/Pointer_Size_10000-12     80357590                15.25 ns/op
PASS
ok      reference-benchmark     13.521
```
