package main

import ("fmt";"math")

type Circle struct {x, y, r float64}

func (im *Circle) Area() float64 {
  return math.Pi * im.r * im.r
}

func main (){

c := Circle {0,0,4}
fmt.Println (c.Area())

}
