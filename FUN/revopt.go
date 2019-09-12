package main

import "fmt"

func rev (sx []int) []int {


for i := len(sx)/2 - 1; i >= 0; i-- {
  opp := len(sx) - 1 - i
  sx[i], sx[opp] = sx[opp], sx[i]

}
return sx
}

func main(){

x := []int{1,2,3,4,5,6}
fmt.Println(x)
fmt.Println(rev(x))
}
