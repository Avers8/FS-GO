package main

import "fmt"

func rev (sx []int) []int {


for i := 0; i < len(sx)/2; i++ {
s1 := sx[i]
s2 := sx[(len(sx) - (i + 1))]
sx[i] = s2
sx[(len(sx) - (i + 1))] = s1

}
return sx
}

func main(){

x := []int{1,2,3,4,5,6}
fmt.Println(x)
fmt.Println(rev(x))
}
