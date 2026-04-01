package main

import (
	"fmt"
)

func main() {
	for {
		var choice string
		fmt.Println("Choose an operation:\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5. Exit")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var a, b float64
			fmt.Println("Enter two numbers to add:")
			fmt.Scanln(&a, &b)
			fmt.Println("Result:", a+b)

		case "2":
			var a, b float64
			fmt.Println("Enter two numbers to subtract:")
			fmt.Scanln(&a, &b)
			fmt.Println("Result:", a-b)

		case "3":
			var a, b float64
			fmt.Println("Enter two numbers to multiply:")
			fmt.Scanln(&a, &b)
			fmt.Println("Result:", a*b)

		case "4":
			var a, b float64
			fmt.Println("Enter two numbers to divide:")
			fmt.Scanln(&a, &b)
			if b == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
			} else {
				fmt.Println("Result:", a/b)
			}

		case "5":
			fmt.Println("Program has ended.")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}