package main

import (
	"fmt"
	"log"
)

func FilterOdd(nums []int) []int {
	//result := []int{}
	var result []int

	for _, num := range nums {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	if len(result) == 0 {
		log.Println("here")
		return nil
	}
	return result
}

func main() {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	var c []int

	fmt.Println("a:", FilterOdd(a)) // → [1 3 5]
	fmt.Println("b:", FilterOdd(b)) // → nil
	fmt.Println("c:", FilterOdd(c)) // → nil

	r := FilterOdd(b)
	if r == nil {
		fmt.Println("Yes, result is nil!")
	} else {
		fmt.Println("Nope, result is not nil :(")
	}
}
