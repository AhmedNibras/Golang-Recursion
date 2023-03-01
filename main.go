package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var printAllSolutions bool
var solutions []string

func parseNumbers(inputSlice []string) ([]int, error) {
	numbers := make([]int, len(inputSlice))
	for i, v := range inputSlice {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		numbers[i] = n
	}
	return numbers, nil
}

func isOperator(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func findSolutions(numbers []int, index int, result int, operation string) {
	if index == len(numbers)-1 {
		solution := operation + " = " + strconv.Itoa(result)
		if result == numbers[index] {
			solutions = append(solutions, solution)
		}
		return
	}

	for _, op := range "+-*/" {
		var newOperation string
		if index == 0 {
			newOperation = strconv.Itoa(numbers[index])
		} else {
			newOperation = operation + " " + string(op) + " " + strconv.Itoa(numbers[index])
		}

		var newResult int
		switch op {
		case '+':
			newResult = result + numbers[index]
		case '-':
			newResult = result - numbers[index]
		case '*':
			newResult = result * numbers[index]
		case '/':
			if numbers[index] != 0 && result%numbers[index] == 0 {
				newResult = result / numbers[index]
			} else {
				continue // Skip division by zero or non-integer results.
			}
		}

		findSolutions(numbers, index+1, newResult, newOperation)
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-all" {
		printAllSolutions = true
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		inputSlice := strings.Fields(input)

		numbers, err := parseNumbers(inputSlice)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing input:", err)
			os.Exit(1)
		}

		solutions = nil
		findSolutions(numbers, 0, 0, "")
		if len(solutions) == 0 {
			fmt.Println()
			continue
		}

		// Remove the first operator in each solution.
		for i := range solutions {
			if isOperator(solutions[i][0]) {
				solutions[i] = solutions[i][1:]
			}
		}

		// Print the solutions.
		if printAllSolutions {
			fmt.Println(strings.Join(solutions, ", "))
		} else {
			fmt.Println(solutions[0])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}
