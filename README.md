## Results

```
goos: darwin
goarch: arm64
pkg: github.com/bcho-msft/rand-string-bench
Benchmark_Generator
Benchmark_Generator/length=8
Benchmark_Generator/length=8/fixed
Benchmark_Generator/length=8/fixed-10   	 7370548	       146.2 ns/op
Benchmark_Generator/length=8/math
Benchmark_Generator/length=8/math-10    	 4750470	       250.2 ns/op
Benchmark_Generator/length=8/crypto
Benchmark_Generator/length=8/crypto-10  	 2411762	       496.1 ns/op
Benchmark_Generator/length=16
Benchmark_Generator/length=16/fixed
Benchmark_Generator/length=16/fixed-10  	 3643945	       328.8 ns/op
Benchmark_Generator/length=16/math
Benchmark_Generator/length=16/math-10   	 2132286	       565.1 ns/op
Benchmark_Generator/length=16/crypto
Benchmark_Generator/length=16/crypto-10 	 2265074	       533.0 ns/op
Benchmark_Generator/length=40
Benchmark_Generator/length=40/fixed
Benchmark_Generator/length=40/fixed-10  	 1357430	       919.9 ns/op
Benchmark_Generator/length=40/math
Benchmark_Generator/length=40/math-10   	  788349	      1494 ns/op
Benchmark_Generator/length=40/crypto
Benchmark_Generator/length=40/crypto-10 	 1901445	       637.7 ns/op
Benchmark_Generator/length=50
Benchmark_Generator/length=50/fixed
Benchmark_Generator/length=50/fixed-10  	 1000000	      1123 ns/op
Benchmark_Generator/length=50/math
Benchmark_Generator/length=50/math-10   	  626770	      1908 ns/op
Benchmark_Generator/length=50/crypto
Benchmark_Generator/length=50/crypto-10 	 1669842	       672.8 ns/op
Benchmark_Generator/length=72
Benchmark_Generator/length=72/fixed
Benchmark_Generator/length=72/fixed-10  	  688668	      1695 ns/op
Benchmark_Generator/length=72/math
Benchmark_Generator/length=72/math-10   	  413032	      2795 ns/op
Benchmark_Generator/length=72/crypto
Benchmark_Generator/length=72/crypto-10 	 1652847	       727.3 ns/op
```


```
goos: linux
goarch: amd64
pkg: github.com/bahe-msft/rand-string-bench
cpu: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz
Benchmark_Generator
Benchmark_Generator/length=8
Benchmark_Generator/length=8/fixed
Benchmark_Generator/length=8/fixed-4   	 3770662	       334.1 ns/op
Benchmark_Generator/length=8/math
Benchmark_Generator/length=8/math-4    	 1767768	       700.0 ns/op
Benchmark_Generator/length=8/crypto
Benchmark_Generator/length=8/crypto-4  	  953038	      1109 ns/op
Benchmark_Generator/length=16
Benchmark_Generator/length=16/fixed
Benchmark_Generator/length=16/fixed-4  	 1602116	       799.0 ns/op
Benchmark_Generator/length=16/math
Benchmark_Generator/length=16/math-4   	  759738	      1498 ns/op
Benchmark_Generator/length=16/crypto
Benchmark_Generator/length=16/crypto-4 	  960842	      1201 ns/op
Benchmark_Generator/length=40
Benchmark_Generator/length=40/fixed
Benchmark_Generator/length=40/fixed-4  	  502593	      2096 ns/op
Benchmark_Generator/length=40/math
Benchmark_Generator/length=40/math-4   	  259202	      4110 ns/op
Benchmark_Generator/length=40/crypto
Benchmark_Generator/length=40/crypto-4 	  640868	      1887 ns/op
Benchmark_Generator/length=50
Benchmark_Generator/length=50/fixed
Benchmark_Generator/length=50/fixed-4  	  406198	      2932 ns/op
Benchmark_Generator/length=50/math
Benchmark_Generator/length=50/math-4   	  228410	      5462 ns/op
Benchmark_Generator/length=50/crypto
Benchmark_Generator/length=50/crypto-4 	  554462	      1934 ns/op
Benchmark_Generator/length=72
Benchmark_Generator/length=72/fixed
Benchmark_Generator/length=72/fixed-4  	  259491	      4247 ns/op
Benchmark_Generator/length=72/math
Benchmark_Generator/length=72/math-4   	  148167	      7983 ns/op
Benchmark_Generator/length=72/crypto
Benchmark_Generator/length=72/crypto-4 	  453092	      2511 ns/op
PASS
```
