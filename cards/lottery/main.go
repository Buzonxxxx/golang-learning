// Self practice:
// random pick up 6 numbers from 1~49
// print out the 6 numbers

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type lottery []int

func main() {
	lotteryNumbers := generateNumbers(49)
	pickedNumbers := pickNumbers(6, lotteryNumbers)
	lotteryNumbers.print()
	pickedNumbers.print()
}

func generateNumbers(endNum int) lottery {
	numbers := []int{}
	for i := 1; i <= endNum; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func pickNumbers(count int, numbers []int) lottery {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	pickNumbers := lottery{}
	for i := 1; i <= count; i++ {
		num := r.Intn(len(numbers))
		pickNumbers = append(pickNumbers, num)
	}
	return pickNumbers
}

func (l lottery) print() {
		fmt.Println(l)
}
