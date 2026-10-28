package main

import "fmt"

func myXor(a int,b int){

	c := a ^ b
	fmt.Println("XOR: ", c)

}

func myNOT(a int ,b int ){

	c := ^(a&b)
  fmt.Println("NOT: ", c)

}

func myOR(a int,b int){

 c := a|b
 fmt.Println("OR: ", c)
 
}

func myAND(a int,b int){

	c := a&b

  fmt.Println("AND: ", c)

}

func leftShift(a int,b int){

	c := a << b

  fmt.Println("left shift: ", c)

}

func rightShift(a int,b int){

	c := a >> b

  fmt.Println("right shift: ", c)

}

func main(){
 

	var num1,num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	 
	myXor(num1,num2)
	myNOT(num1,num2)
	myOR(num1,num2)
	myAND(num1,num2)
	leftShift(num1,num2)
	rightShift(num1,num2)

}
