package main

import "fmt"

func fibonachi(n uint) uint{


if n < 2{
  return n
}

return fibonachi(n-2) + fibonachi(n-1)

}
func main() {

    fmt.Println(fibonachi(7))


}
