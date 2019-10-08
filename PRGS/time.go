package main

import "fmt"

func gradus (h,m float64)  float64 {

r := (h * 30 + m / 2) - m * 6

if r < 0 {
   return - r
}

return r

}

func main() {
  h := 1.0
  m := 10.0

  fmt.Println(gradus(h,m))
}
