package main

import "fmt"

func f(x int, y int) (int, int) {

	return x + y, 7
}

func main() {

	x1 := 2
	x2 := 4

	res, n := f(x1, x2)
	fmt.Println(res, n)

}
