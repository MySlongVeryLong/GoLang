package main

import "fmt"

func main(){
 

	var num1,num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)


	fmt.Println("num1 && num2: ", num1 > 0 && num2 > 0) 
	fmt.Println("num1 > num2: ", num1 > num2)
	fmt.Println("num1 != num2: ", num1 != num2)



}
