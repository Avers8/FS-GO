package main

import "fmt"

func f(x int, y int) int {

z := 0
z = x + y
return z
}

func main () {

x1 := 2
x2 := 4

res := f(x1,x2)
fmt.Println (res)

}
