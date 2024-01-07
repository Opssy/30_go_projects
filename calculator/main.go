package main

import "fmt"



func main(){
  var operator string
  var a, b int

   fmt.Println("Please enter first number: ")
   fmt.Scan(& a)

   fmt.Println("Please enter first number: ")
   fmt.Scan(& b)

    fmt.Println("Please enter operator(+,_,/,*,%):")
   fmt.Scan(& operator)

    result := 0 

  switch operator {

    case "+":
      result = a + b
    break

    case "-":
      result = a - b
    break

    case "*":
      result = a * b
    break

    case "/":
      result = a / b
    break

    case "%":
      result = a % b
    break

    default:
    fmt.Println("Invalid Operation")
  }
  fmt.Printf("%d %s %d = %d", a, operator, b, result)
}