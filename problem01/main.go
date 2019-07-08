package main

import (
	"daily-coding-problem/utilities"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// #1 Google
// Given a list of numbers and a number k, return whether any two numbers from the list add up to k
// For example, given [10, 15, 3, 7] and k of 17 return true since 10 + 7 is 17

func main() {
	fmt.Println("Enter a list of numbers separated by commas")
	fmt.Println("Example: 1,59,20,234,20\n")

	fmt.Print("-> ")

	input := utilities.ReadConsoleInput()

	sanitizedInput := utilities.SanitizeInput(input)

	trimmedInput := strings.Split(sanitizedInput, ",")

	numbers, err := utilities.ListOfNumbersToInt(trimmedInput)

	if err != nil {
		log.Fatal(err.Error())
	}

	totalNumbers := len(numbers)

	fmt.Println()
	fmt.Println("Enter a number")

	fmt.Print("-> ")

	input = utilities.ReadConsoleInput()

	sanitizedInput = utilities.SanitizeInput(input)

	number, _ := strconv.Atoi(sanitizedInput)

	flag := false

	for i := 0; i <= totalNumbers-1; i++ {
		for j := i + 1; j < totalNumbers; j++ {
			if numbers[i]+numbers[j] == number {
				flag = true
				break
			}
		}
	}

	fmt.Println(flag)
}
