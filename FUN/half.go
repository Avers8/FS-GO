package main

import "fmt"

func hf(n int) (int, bool) {
return n/2, n%2 == 0
}

func main () {
fmt.Println (hf(372))
}
