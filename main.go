package main

import (
	"fmt"

	"calculater/calc"
)

func main() {

	var num1, num2 int
	var choice string

	fmt.Println("Enter first num")
	fmt.Scan(&num1)

	fmt.Println("Enter your operater")
	fmt.Scan(&choice)

	fmt.Println("Enter second num")
	fmt.Scan(&num2)

	switch choice {
	case "+":
		fmt.Printf("Addition: %d \n", calc.Addition(num1, num2))
	case "-":
		fmt.Printf("Difference: %d \n", calc.Substraction(num1, num2))
	case "*":
		fmt.Printf("Product: %d \n", calc.Multiplication(num1, num2))
	case "/":
		fmt.Printf("Division: %d \n", calc.Division(num1, num2))
	default:
		fmt.Println("Wrong option try with +, -, / and *")
	}

}
