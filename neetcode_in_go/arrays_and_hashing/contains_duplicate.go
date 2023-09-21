package main

import (
	"fmt"
	"neetcodeingo/utils"
)

func containsDuplicate(nums []int) bool {
	ref := make(map[int]bool)

	for _, n := range nums {
		if ref[n] {
			return true
		}
		ref[n] = true
	}

	return false
}

func main() {
	nums := utils.ReadNumbersFromInput(", ")

	fmt.Println(containsDuplicate(nums))
}
