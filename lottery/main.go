// Self practice:
// random pick up 6 numbers from 1~49
// print out the 6 numbers

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	lotteryNumbers := generateNumbers(49)
	pickNumbers := pickNumbers(6, lotteryNumbers)
	print(pickNumbers)
}

func generateNumbers(endNum int) []int {
	numbers := []int{}
	for i := 1; i <= endNum; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func pickNumbers(count int, numbers []int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	pickNumbers := []int{}
	for i := 1; i <= count; i++ {
		num := r.Intn(len(numbers))
		pickNumbers = append(pickNumbers, num)
	}
	return (pickNumbers)
}

func print(numbers []int) {
	fmt.Println(numbers)
}
