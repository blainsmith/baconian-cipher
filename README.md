# Baconian cipher

[Baconian cipher](https://en.wikipedia.org/wiki/Bacon%27s_cipher) in Go.

## Benchmarks
```
$ go test -bench . -benchmem
BenchmarkEncrypt-4   	  300000	      4997 ns/op	    1808 B/op	      14 allocs/op
BenchmarkDecrypt-4   	  500000	      3871 ns/op	     672 B/op	      11 allocs/op
```
