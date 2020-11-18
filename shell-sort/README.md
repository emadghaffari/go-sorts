# sell sort: (for very large lists)
## idea from insertion sort
## O(nlogn) === n scan list, logn passes
## O(n power 3/2)
## cap=[n/2], n = list length

test result:
#### BenchmarkInsertion-8   	   40254	     27115 ns/op	    2328 B/op	      20 allocs/op
#### BenchmarkInsertion-8   	   42352	     27455 ns/op	    2328 B/op	      20 allocs/op
#### BenchmarkInsertion-8   	   42352	     27455 ns/op	    2328 B/op	      20 allocs/op
#### BenchmarkInsertion-8   	   48888	     26667 ns/op	    2328 B/op	      20 allocs/op
