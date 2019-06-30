package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a list of numbers separated by commas")
	fmt.Println("Example: 1,23,50,12")

	fmt.Print("-> ")

	scanner.Scan()
	input := scanner.Text()
	trimmedInput := strings.TrimSpace(input)
	numbersStrings := strings.Split(trimmedInput, ",")

	numbers := make([]int, len(numbersStrings))

	for index, number := range numbersStrings {
		numbers[index], _ = strconv.Atoi(number)
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
