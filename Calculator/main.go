package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Print("Enter operation (add, substract, multiply, divide): ")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number:")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter second number: ")
	fmt.Scanf("%d", &num2)

	switch operation {
	case "add":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "substract":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "multiply":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "divide":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	default:
		fmt.Printf("Unknown operation")
	}
}
