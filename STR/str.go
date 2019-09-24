package main

import ("fmt") //;"math")

type Circle struct {x, y, r float64}

func Change (im *Circle) {
    *im.r = 12
                      }

func main (){

c := Circle{0,0,5}

Change(&c)

fmt.Println(c)
}
