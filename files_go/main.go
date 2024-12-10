package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main is the entry point of the program. It presents a menu to the user with five options:
// 1. Number Crunch
// 2. Players
// 3. Score
// 4. Count Letters
// 5. Exit
//
// The user is prompted to enter a number corresponding to one of the options. Based on the user's choice,
// the appropriate function is called. If the user enters an invalid input or an invalid choice, an error
// message is displayed and the menu is shown again.
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Number Crunch")
		fmt.Println("2. Players")
		fmt.Println("3. Score")
		fmt.Println("4. Count Letters")
		fmt.Println("5. Exit")
		fmt.Print("Enter an option: ")

		var choice int
		_, err := fmt.Fscanf(reader, "%d\n", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			reader.ReadString('\n') // Clear the buffer
			continue
		}

		switch choice {
		case 1:
			numCrunch(reader)
		case 2:
			players(reader)
		case 3:
			score(reader)
		case 4:
			countLetters(reader)
		case 5:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

// numCrunch performs a basic arithmetic operation (+, -, *, /) on two numbers
// provided by the user through the given bufio.Reader. The result of the
// operation is written to a file named "calculation.txt" and also printed to
// the console.
//
// The function prompts the user to choose an operation and enter two numbers.
// It validates the input and handles errors such as invalid operations,
// invalid numbers, and division by zero.
//
// Parameters:
//
//	reader (*bufio.Reader): A buffered reader to read user input.
//
// Example usage:
//
//	reader := bufio.NewReader(os.Stdin)
//	numCrunch(reader)
func numCrunch(reader *bufio.Reader) {
	file, err := os.Create("calculation.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Print("Choose an operation (+, -, *, /): ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		fmt.Println("Invalid operation")
		return
	}

	fmt.Print("Enter a number: ")
	var num1 float64
	_, err = fmt.Fscanf(reader, "%f\n", &num1)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	fmt.Print("Enter another number: ")
	var num2 float64
	_, err = fmt.Fscanf(reader, "%f\n", &num2)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		result = num1 / num2
	}

	file.WriteString(fmt.Sprintf("%.2f\n", result))
	fmt.Println("The result is:", result)
}

// players prompts the user to enter the names of four players, writes these names to a file named "players2.txt",
// and then reads and prints the contents of the file. If there is an error creating or reading the file, it prints an error message.
//
// Parameters:
//
//	reader: a bufio.Reader used to read user input from the console.
func players(reader *bufio.Reader) {
	file, err := os.Create("players2.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 0; i < 4; i++ {
		fmt.Print("Enter a player name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		file.WriteString(name + "\n")
	}

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// score reads a score from the provided bufio.Reader and appends it to a file named "scores.txt".
// If the file does not exist, it will be created. If there is an error opening the file or reading
// the input, an error message will be printed to the console.
//
// Parameters:
//
//	reader: A pointer to a bufio.Reader from which the score will be read.
//
// The function expects the input to be a valid integer. If the input is not a valid integer,
// an error message will be printed and the function will return without writing to the file.
func score(reader *bufio.Reader) {
	file, err := os.OpenFile("scores.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Enter your latest score:")
	var score int
	_, err = fmt.Fscanf(reader, "%d\n", &score)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	fmt.Fprintf(file, "%d\n", score)
}

// countLetters reads a letter from the user, counts its occurrences in the file "romeojuliet.txt",
// prints the count to the console, and appends the letter and its count to the file.
//
// Parameters:
//
//	reader (*bufio.Reader): A buffered reader to read user input.
//
// The function performs the following steps:
// 1. Prompts the user to enter a letter.
// 2. Reads the letter from the user input and trims any surrounding whitespace.
// 3. Checks if the input is a single letter; if not, prints an error message and returns.
// 4. Reads the content of the file "romeojuliet.txt".
// 5. Counts the occurrences of the letter in the file content.
// 6. Prints the count of the letter to the console.
// 7. Opens the file "romeojuliet.txt" in append mode.
// 8. Appends the letter and its count to the file.
func countLetters(reader *bufio.Reader) {
	fmt.Print("Enter a letter to count: ")
	letter, _ := reader.ReadString('\n')
	letter = strings.TrimSpace(letter)
	if len(letter) != 1 {
		fmt.Println("Please enter only one letter")
		return
	}

	content, err := os.ReadFile("romeojuliet.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	count := strings.Count(string(content), letter)

	fmt.Printf("The letter '%s' appears %d times\n", letter, count)

	file, err := os.OpenFile("romeojuliet.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "\n\nLetter: %s\nCount: %d\n", letter, count)
}
