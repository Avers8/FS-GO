package main

import "fmt"




func makeEvenGenerator() func() uint { // возврат функции и целого числа
  i := uint(0)   //переменная внешней функции
  return func() (ret uint) {  // внутренная анонимная функция возвращает ret
    ret = i   // доступ к i из внешней функции
    i += 2  // изменение i
    return // в итоге возврат ret и изменение i
  }
}
func main() {
  nextEven := makeEvenGenerator() // в переменной находится функция
  fmt.Println(nextEven()) // 0  // вызывается внутренняя функция с доступом к i
  fmt.Println(nextEven()) // 2  // i изменился и вывелся внутренней функцией
  fmt.Println(nextEven()) // 4
}
