package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// #1 Google
// Given a list of numbers and a number k, return whether any two numbers from the list add up to k
// For example, given [10, 15, 3, 7] and k of 17 return true since 10 + 7 is 17

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a list of numbers separated by commas")
	fmt.Println("Example: 1,59,20,234,20\n")

	fmt.Print("-> ")

	scanner.Scan()
	input := scanner.Text()
	inputTrimmed := strings.TrimSpace(input)

	numbersStrings := strings.Split(inputTrimmed, ",")
	totalNumbers := len(numbersStrings)

	numbers := make([]int, totalNumbers)

	for index, number := range numbersStrings {
		numbers[index], _ = strconv.Atoi(number)
	}

	fmt.Println()
	fmt.Println("Enter a number")

	fmt.Print("-> ")

	scanner.Scan()
	input = scanner.Text()
	inputTrimmed = strings.TrimSpace(input)

	number, _ := strconv.Atoi(inputTrimmed)

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
