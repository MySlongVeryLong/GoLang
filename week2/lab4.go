package main

import "fmt"

func add(a int,b int) int {

	c := a + b
	return c

}

func sub(a int ,b int) int{

	c := a - b
  return c

}

func mul(a int,b int) int{

 c := a * b
 return c
 
}

func div(a int,b int) int {
  
	if(b == 0){
		fmt.Println("Error: Division by zero")
		return 0
	}
	c := a / b
	return c

}

func mod(a int,b int) int{

	c := a % b
	return c

}


func main(){
 
	var a,b,option int
  
	fmt.Print(`

	===== Mini Calculator =====
  1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit

	`)

  fmt.Print("Choose: ")
	fmt.Scan(&option)
  if(option == 6){
		fmt.Println("Exiting program")
		return
	}

	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)

  // switch option
	switch option {
	case 1:
		fmt.Println("Add: ",add(a,b))
	case 2:
		fmt.Println("Sub: ",sub(a,b))
	case 3:
		fmt.Println("Mul: ", mul(a,b))
	case 4:
		fmt.Println("Div:", div(a,b))
	case 5:
		fmt.Println("Mod:", mod(a,b))
	default:
			fmt.Println("Invalid option")

	}

}
