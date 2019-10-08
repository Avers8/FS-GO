package main

import (
	"fmt"; "time"

)

func printHello() {
  fmt.Println("Hello World")
}

func main() {
  fmt.Println("main execution started")

//call function
go printHello()

time.Sleep(10 * time.Millisecond)

fmt.Println("main execution stopped")
}
