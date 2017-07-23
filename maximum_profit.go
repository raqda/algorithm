package main

import "fmt"

func main() {
    var n,min,r int
    fmt.Scan(&n)
    fmt.Scan(&min)
    fmt.Scan(&r)
    var max int = r-min;

    for i := 2; i < n; i++ {
	    var r int;
	    fmt.Scan(&r)

	    if max < r-min {
	        max = r-min;
	    }
	    if min > r {
	        min = r;
	    }
    }
    fmt.Printf("%d\n", max)
}
