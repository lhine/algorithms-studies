package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumbersFromInput(splitStr string) []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	numbersStr := strings.Split(input, splitStr)
	numbers := make([]int, len(numbersStr))

	for i, numberStr := range numbersStr {
		num, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println("error:", err)
			return nil
		}
		numbers[i] = num
	}

	return numbers
}
