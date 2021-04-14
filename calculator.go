package main

import "fmt"

func main() {
	var operator string
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	fmt.Print("Select operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	output := 0
	switch operator {
	case "+":
		output = num1 + num2
		break
	case "-":
		output = num1 - num2
	case "*":
		output = num1 * num2
	case "/":
		output = num1 / num2
	default:
		fmt.Println("Invalid Operator!")
	}
	fmt.Printf("%d %s %d = %d", num1, operator, num2, output)
}
