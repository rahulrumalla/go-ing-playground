package leetcode

import (
	"strconv"
)

//https://leetcode.com/problems/reverse-integer/description/
// Reverse digits of an integer.
// Example1: x = 123, return 321
// Example2: x = -123, return -321
// The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.

func ReverseInteger(x int) int {
	if x == 0 {
		return 0
	}

	isNeg := x < 0
	xString := strconv.Itoa(x)

	var xStringReverse string
	if isNeg {
		xStringReverse = Reverse(xString[1:])
	} else {
		xStringReverse = Reverse(xString)
	}

	i64, err := strconv.ParseInt(xStringReverse, 10, 32)
	if err != nil {
		return 0
	}

	if isNeg {
		i64 *= -1
	}

	return int(i64)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
