
# inserion sort: good for sort linkless!
## compares min: O(n)
## compares max: O(n²)
## swap min: (n)
## swap max: n(n-1)/2 === O(n²)

test result:

#### BenchmarkInsertion-8   	   56042	     21116 ns/op	    2040 B/op	       8 allocs/op
#### BenchmarkInsertion-8   	   55264	     21462 ns/op	    2040 B/op	       8 allocs/op


insertion-sort is adaptive