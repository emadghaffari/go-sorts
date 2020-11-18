
# selection sort: good for less numbers swaps (only algorithm with less number swap)
## is NOT Adaptive or Stable
## compares min: O(n)
## compares max: O(n²)
## swap min: (n)
## swap max: n(n-1)/2 === O(n²)


#### BenchmarkSelectionSort-8   	   50127	     22877 ns/op	    2040 B/op	       8 allocs/op
#### BenchmarkSelectionSort-8   	   51513	     23491 ns/op	    2040 B/op	       8 allocs/op
#### BenchmarkSelectionSort-8   	   50432	     24454 ns/op	    2040 B/op	       8 allocs/op
