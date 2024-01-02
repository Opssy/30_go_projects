package main

import "fmt"



func add(a, b int) int {
  return a + b
}

func sub(a,b int) int {
	return a - b
}
func mul(a,b int) int {
	return a * b
}

func div(a,b int) int {
   if  b == 0 {
	fmt.Println("cannot be divided by zero")
   }
   return a / b
}


func main(){
   fmt.Println(add, "Enter your number:")
   
}