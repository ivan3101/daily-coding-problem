package main

import "daily-coding-problem/utilities"
import "fmt"
import "log"
import "strings"

func main() {
	fmt.Println("Enter a list of numbers separated by commas")
	fmt.Println("Example: 1,23,50,12")

	fmt.Print("-> ")

	input := utilities.ReadConsoleInput()

	sanitizedInput := utilities.SanitizeInput(input)

	splittedInput := strings.Split(sanitizedInput, ",")

	numbers, err := utilities.ListOfNumbersToInt(splittedInput)

	if err != nil {
		log.Fatal(err.Error())
	}

	numbersProducts := make([]int, len(numbers))

	for index, _ := range numbers {
		result := 1
		for index2, number := range numbers {
			if index != index2 {
				result = result * number
			}
		}
		numbersProducts[index] = result
	}

	fmt.Println(numbersProducts)
}
