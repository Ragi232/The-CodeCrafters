package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var operations string
		fmt.Println("Select an operation to continue...\n 1.Hex\n 2.Bin\n 3.Decimal")
		fmt.Scanln(&operations)

		switch operations {
		case "1":
			{
				var input string
				fmt.Println("Input the values to be converted")
				fmt.Scanln(&input)

				decimal, err := strconv.ParseInt(input, 16, 64)
				if err != nil {
					fmt.Println("Invalid hexadecimal input... Try again")
					continue
				}
				
				fmt.Println("Result:", decimal)
			}
		case "2":
			{
				var input string
				fmt.Println("Enter the values to be converted")
				fmt.Scanln(&input)

				decimal, err := strconv.ParseInt(input, 2, 64)
				if err != nil {
					fmt.Println("Invalid Binary number input.. Try again")
					continue
				}
				fmt.Println("Result:", decimal)
			}
		case "3":
			{
				var input string
				fmt.Println("Input the values to be converted")
				fmt.Scanln(&input)

				decimal, err := strconv.ParseInt(input, 10, 64)
				if err != nil {
					fmt.Println("Invalid value.. Try again")
					return
				}
				binary := strconv.FormatInt(decimal, 2)
				hexadecimal := strconv.FormatInt(decimal, 16)

				fmt.Println("Result (Binary):", binary)
				fmt.Println("Result (hexadeccimal):", strings.ToUpper(hexadecimal))

				return
			}
		}
	}
}