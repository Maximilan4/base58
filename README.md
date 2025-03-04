# base58
Go wrap of https://github.com/moonlibs/base58/blob/master/libbase58.c created for [benchmark](https://github.com/Maximilan4/base58_benchmark_comparison)
```
Note: This is slow. Use better alternative instead.
```
```shell
go get github.com/Maximilan4/base58
```


## Encode
```go 
// encode raw bytes
src := []byte("09")
dst := make([]byte, 10)
n, err := base58.EncodeBytes(src, dst)
if err != nil {
	log.Fatal(err)
}
fmt.Println(dst[:n])

// or encode string 
trg, err := base58.EncodeString("09")
if err != nil {
	log.Fatal(err)
}
fmt.Println(trg)

```

## Decode
```go 
src := []byte("4ER")
dst := make([]byte, 10)
n, err := base58.DecodeBytes(src, dst)
if err != nil {
    log.Fatal(err)
}
fmt.Println(dst[:n])

// or encode string 
trg, err := base58.DecodeString("4ER")
if err != nil {
    log.Fatal(err)
}
fmt.Println(trg)
```

### Benchmarking
```shell
# source: random
# cases: power or 2
# bytes encoding 
goos: darwin
goarch: amd64
pkg: github.com/Maximilan4/base58
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkEncodeBytes/0-16      592855269               1.948 ns/op           0 B/op          0 allocs/op
BenchmarkEncodeBytes/2-16        9518738               113.9 ns/op           8 B/op          1 allocs/op
BenchmarkEncodeBytes/4-16        8279209               138.9 ns/op           8 B/op          1 allocs/op
BenchmarkEncodeBytes/8-16        5676512               211.0 ns/op           8 B/op          1 allocs/op
BenchmarkEncodeBytes/16-16       2337645               506.8 ns/op           8 B/op          1 allocs/op
BenchmarkEncodeBytes/32-16        590566               1750 ns/op            8 B/op          1 allocs/op
BenchmarkEncodeBytes/64-16        151405               6913 ns/op            8 B/op          1 allocs/op
BenchmarkEncodeBytes/128-16        42712               27531 ns/op           8 B/op          1 allocs/op
BenchmarkEncodeBytes/256-16        10000               114411 ns/op          8 B/op          1 allocs/op
BenchmarkEncodeBytes/512-16         2628               434018 ns/op          8 B/op          1 allocs/op
BenchmarkEncodeBytes/1024-16         676               1706298 ns/op         8 B/op          1 allocs/op
BenchmarkEncodeBytes/2048-16         172               6914281 ns/op         8 B/op          1 allocs/op
PASS
ok      github.com/Maximilan4/base58    16.533s

# string encoding 
goos: darwin
goarch: amd64
pkg: github.com/Maximilan4/base58
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkEncodeString/0-16      668270265            1.726 ns/op             0 B/op           0 allocs/op
BenchmarkEncodeString/2-16      10677853             105.2 ns/op             16 B/op          2 allocs/op
BenchmarkEncodeString/4-16       9039418             126.6 ns/op             16 B/op          2 allocs/op
BenchmarkEncodeString/8-16       5690422             205.4 ns/op             24 B/op          2 allocs/op
BenchmarkEncodeString/16-16      2356365             575.0 ns/op             40 B/op          2 allocs/op
BenchmarkEncodeString/32-16       607369              1887 ns/op             72 B/op          2 allocs/op
BenchmarkEncodeString/64-16       158307              7666 ns/op             120 B/op         2 allocs/op
BenchmarkEncodeString/128-16       42396             27209 ns/op             216 B/op         2 allocs/op
BenchmarkEncodeString/256-16        9960            115543 ns/op             424 B/op         2 allocs/op
BenchmarkEncodeString/512-16        2245            511993 ns/op             904 B/op         2 allocs/op
BenchmarkEncodeString/1024-16        548           1896499 ns/op            1800 B/op         2 allocs/op
BenchmarkEncodeString/2048-16        140           7237347 ns/op            3208 B/op         2 allocs/op
PASS
ok      github.com/Maximilan4/base58    17.887s

# bytes decoding (encoded random data)
goos: darwin
goarch: amd64
pkg: github.com/Maximilan4/base58
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecodeBytes/0-16       519316422            2.360 ns/op           0 B/op          0 allocs/op
BenchmarkDecodeBytes/3-16        9494430             121.8 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/6-16        8508218             137.1 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/11-16       6758929             169.0 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/22-16       4224823             282.7 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/44-16       1989705             588.9 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/88-16        670935              1892 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/175-16       148191              6922 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/350-16        43537             29384 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/699-16         9264            124445 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/1398-16        2264            487703 ns/op           8 B/op          1 allocs/op
BenchmarkDecodeBytes/2797-16         552           2122545 ns/op           8 B/op          1 allocs/op

# string decoding (encoded random data)
goos: darwin
goarch: amd64
pkg: github.com/Maximilan4/base58
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecodeString/0-16      549042093            2.143 ns/op               0 B/op          0 allocs/op
BenchmarkDecodeString/3-16       8667056             142.9 ns/op              16 B/op          2 allocs/op
BenchmarkDecodeString/6-16       7339426             160.2 ns/op              24 B/op          2 allocs/op
BenchmarkDecodeString/11-16      6231201             185.7 ns/op              24 B/op          2 allocs/op
BenchmarkDecodeString/22-16      4033837             304.9 ns/op              40 B/op          2 allocs/op
BenchmarkDecodeString/44-16      1896812             558.7 ns/op              72 B/op          2 allocs/op
BenchmarkDecodeString/88-16       612112              1863 ns/op             136 B/op          2 allocs/op
BenchmarkDecodeString/175-16      157795              6792 ns/op             248 B/op          2 allocs/op
BenchmarkDecodeString/349-16       44415             27487 ns/op             488 B/op          2 allocs/op
BenchmarkDecodeString/700-16       10000            114872 ns/op            1032 B/op          2 allocs/op
BenchmarkDecodeString/1399-16       2445            440904 ns/op            2056 B/op          2 allocs/op
BenchmarkDecodeString/2797-16        664           1753002 ns/op            4104 B/op          2 allocs/op
PASS
ok      github.com/Maximilan4/base58    17.165s
```
