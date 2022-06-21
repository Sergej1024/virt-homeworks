package main

import "fmt"

func main() {
    var n int
    var m int
    n = 1
    m = 100

    for i := n; i <= m; i++ {
	if i%3 == 0 {
	    fmt.Print(i, " ")
	}
    }
}