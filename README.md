<p align="center">
  <b>
    <span style="font-size:larger;">fastgcd-go</span>
  </b>
  <br />
   <a href="https://travis-ci.org/detailyang/fastgcd-go"><img src="https://travis-ci.org/detailyang/fastgcd-go.svg?branch=master" /></a>
   <br />
   <b>fastgcd-go implements the binary gcd in go</b>
</p>


````golang
╰─λ go test -benchtime 10s -benchmem -bench "." .
goos: darwin
goarch: amd64
pkg: github.com/detailyang/fastgcd-go
BenchmarkVanillaGCD-8   	200000000	       100 ns/op	      16 B/op	       2 allocs/op
BenchmarkGCD-8          	500000000	        31.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinaryGCD-8    	2000000000	        10.5 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/detailyang/fastgcd-go	75.654s
````
