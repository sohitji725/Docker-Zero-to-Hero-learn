package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi Abhishek.Veeramalla, I am a calculator app ....")

	for {
		// Read input from the user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter any calculation (Example: 1 + 2 or 2 * 5) -> Please maintain spaces as shown: ")
		text, _ := reader.ReadString('\n')

		// Trim the newline character from the input
		text = strings.TrimSpace(text)

		// Exit condition
		if text == "exit" {
			fmt.Println("Exiting calculator...")
			break
		}

		// Split input into parts
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Convert operands to integers
		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[2])
		operator := parts[1]

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid numbers. Try again.")
			continue
		}

		// Perform the operation
		var result int
		switch operator {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				fmt.Println("Cannot divide by zero.")
				continue
			}
			result = left / right
		default:
			fmt.Println("Unsupported operator. Use +, -, *, or /")
			continue
		}

		// Print the result
		fmt.Printf("Result: %d\n", result)
	}
}
