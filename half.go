package main

import "fmt"

func hf(n int) (half int, bul bool) {
res := n / 2
if res % 2 == 0 {
  bul = true
} else {
  bul = false
}
return res, bul
}

func main () {
x := 4
fmt.Println (hf(x))
}
