package main

import "fmt"

func f(x int, y int) (p int) {

z := 0
z = x + y
p = z
return
}

func main () {

x1 := 2
x2 := 4

res := f(x1,x2)
fmt.Println (res)

}
