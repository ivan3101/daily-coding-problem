package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ListOfNumbersToInt(listNumbers []string) ([]int, error) {
	numbers := make([]int, len(listNumbers))

	var err error

	for index, number := range listNumbers {
		numbers[index], err = strconv.Atoi(number)

		if err != nil {
			return nil, err
		}
	}

	return numbers, nil
}

func ReadConsoleInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}

func SanitizeInput(input string) string {
	trimmedInput := strings.TrimSpace(input)

	return trimmedInput
}
