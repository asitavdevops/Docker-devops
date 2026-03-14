package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printHelp() {
	fmt.Println("\nAvailable operations:")
	fmt.Println("  +  Addition")
	fmt.Println("  -  Subtraction")
	fmt.Println("  *  Multiplication")
	fmt.Println("  /  Division")
	fmt.Println("\nExample: 10 + 5")
	fmt.Println("Type 'exit' to quit the calculator")
	fmt.Println("Type 'help' to display this help message\n")
}

func main() {

	fmt.Println("========================================")
	fmt.Println("Hi Asitav.Pattanaik 👋")
	fmt.Println("Welcome to the Go CLI Calculator App")
	fmt.Println("========================================")

	printHelp()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter calculation: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Try again.")
			continue
		}

		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Thank you for using the calculator. Goodbye!")
			break
		}

		if text == "help" {
			printHelp()
			continue
		}

		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid format. Example: 5 + 3")
			continue
		}

		left, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid first number")
			continue
		}

		right, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			fmt.Println("Invalid second number")
			continue
		}

		operator := parts[1]

		var result float64

		switch operator {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				fmt.Println("Error: Division by zero is not allowed")
				continue
			}
			result = left / right
		default:
			fmt.Println("Unsupported operator. Use + - * /")
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}
}
