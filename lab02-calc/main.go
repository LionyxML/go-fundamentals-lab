package main

import "fmt"

func main() {
	var option string
	var num1 float64
	var num2 float64

	fmt.Println("WELCOME TO GO CALC!")
	fmt.Println("===================")
	fmt.Println("Please select one operation (add, sub, mult, div):")
	fmt.Scanf("%s", &option)
	fmt.Println("Enter the first number:")
	fmt.Scanf("%v", &num1)
	fmt.Println("Enter the second number:")
	fmt.Scanf("%v", &num2)
	switch option {
	case "add":
		fmt.Printf("%v + %v = %v", num1, num2, num1+num2)
	case "sub":
		fmt.Printf("%v - %v = %v", num1, num2, num1-num2)
	case "mult":
		fmt.Printf("%v * %v = %v", num1, num2, num1*num2)
	case "div":
		fmt.Printf("%v / %v = %v", num1, num2, num1/num2)
	default:
		fmt.Printf("oops, can't perform this calculation")
		panic("AHHHHHHHHHHHHHHHHHHH...")
	}
}
