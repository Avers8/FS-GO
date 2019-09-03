package main

import "fmt"


func sum(n []int) int {

total := 0
for _, v := range n {
  total += v
  }
  return total

}

func main() {

res := []int{3,43,22,67,99}

fmt.Println(sum(res))
}
