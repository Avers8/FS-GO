package main

import "fmt"

func add(args ...int) int {

res := 0

for _, v := range args  {
    if v > res {
     res = v }

  }
return res
}

func main() {
  xs := []int{1,345,4,5,45}
  fmt.Println(add(xs...))
}
