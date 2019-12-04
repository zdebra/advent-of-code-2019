package main

import (
	"fmt"
	"strconv"
)

type testFunc func(int) bool

// At least two adjacent digits are the same (like 22 in 122345)
func adjacentDigits(n int) bool {
	nStr := strconv.Itoa(n)
	prev := nStr[0]
	for i := 1; i < len(nStr); i++ {
		if nStr[i] == prev {
			return true
		}
		prev = nStr[i]
	}
	return false
}

// Going from left to right, the digits never decrease;
// they only ever increase or stay the same (like 111123 or 135679)
// it's turner around
func upwardTrend(n int) bool {
	prevNumber := n % 10
	n = n / 10
	for n > 0 {
		digit := n % 10
		if digit > prevNumber {
			return false
		}
		prevNumber = digit
		n = n / 10
	}
	return true
}

func matchCriteria(n int) bool {
	tests := []testFunc{adjacentDigits, upwardTrend}
	for _, tFunc := range tests {
		if !tFunc(n) {
			return false
		}
	}
	return true
}

func main() {
	total := 0
	for i := 156218; i <= 652527; i++ {
		if matchCriteria(i) {
			total++
		}
	}
	fmt.Println(total)
}
