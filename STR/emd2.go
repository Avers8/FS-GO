package main

import ("fmt")

type ball struct {
  radius int
  material string
}

type football struct {
  ball
}

func (b ball) bounce() {
  fmt.Println(b)
}

func main(){
  fb := football{ball{radius: 5, material: "leather"}}
  fb.ball.bounce()
}
