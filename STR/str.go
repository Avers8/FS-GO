package main

import ("fmt") //;"math")

type Circle struct {x, y, r float64}

func Change (im *float64) {
    *im = 34
                      }

func main (){

c := Circle{0,0,5}

Change(&c.r)

fmt.Println(c)
}
