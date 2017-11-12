package leetcode

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseInteger(t *testing.T) {
	input := 123
	expected := 321
	actual := ReverseInteger(input)
	assert.Equal(t, expected, actual, "Not same")

	input = -123
	expected = -321
	actual = ReverseInteger(input)
	assert.Equal(t, expected, actual, "Not same")

	input = math.MaxInt64
	expected = 0
	actual = ReverseInteger(input)
	assert.Equal(t, expected, actual, "Not same")
}

func TestTwoSum(t *testing.T) {
	inputs := []int{3, 3}
	target := 6
	expected := []int{0, 1}
	actual := TwoSum(inputs, target)

	assert.Equal(t, expected, actual, "Not Same")
}

func TestSimplifyPath(t *testing.T) {
	input := "/home/./b/"
	expect := "/home/b"
	actual := SimplifyPath(input)
	assert.Equal(t, expect, actual, "not same")

	input = "/a/./b/../c/"
	expect = "/a/c"
	actual = SimplifyPath(input)
	assert.Equal(t, expect, actual, "not same")

	input = "/a/./b/../../c/"
	expect = "/c"
	actual = SimplifyPath(input)
	assert.Equal(t, expect, actual, "not same")

	input = "/home//foo"
	expect = "/home/foo"
	actual = SimplifyPath(input)
	assert.Equal(t, expect, actual, "not same")

	input = "/."
	expect = "/"
	actual = SimplifyPath(input)
	assert.Equal(t, expect, actual, "not same")

	input = "/"
	expect = "/"
	actual = SimplifyPath(input)
	assert.Equal(t, expect, actual, "not same")

	input = "/home../../.."
	expect = "/"
	actual = SimplifyPath(input)
	assert.Equal(t, expect, actual, "not same")
}

func TestStringSlice(t *testing.T) {
	input := "hello"
	firstChar := input[:1]
	assert.Equal(t, firstChar, "h", "Not same")

	lastChar := input[len(input)-1:]
	assert.Equal(t, lastChar, "o", "Not same")

	hell := input[:len(input)-1]
	assert.Equal(t, hell, "hell", "Not same")

	e := input[1:2]
	assert.Equal(t, e, "e", "Not same")
}

func TestTriangleNumber(t *testing.T) {
	TriangleNumber([]int{66, 77, 91, 38, 14, 86, 73, 9, 64, 38, 61, 58, 87, 41, 10, 73, 57, 12, 67, 26, 87, 49, 97, 1, 71, 24, 95, 27, 16, 75, 99, 59, 77, 74, 26, 19, 45, 86, 48, 60, 61, 30, 90, 11, 94, 36, 46, 16, 72, 16, 80, 35, 25, 62, 64, 11, 78, 65, 71, 98, 21, 58, 89, 50, 0, 96, 41, 95, 25, 24, 84, 60, 94, 97, 42, 35, 30, 21, 44, 33, 41, 35, 83, 7, 35, 64, 20, 93, 75, 10, 91, 76, 58, 94, 25, 46, 83, 75, 69, 83})
}

func TestAddTwoNumbers(t *testing.T) {
	l3 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 4, Next: l3}
	l1 := &ListNode{Val: 2, Next: l2}

	r3 := &ListNode{Val: 4}
	r2 := &ListNode{Val: 5, Next: r3}
	r1 := &ListNode{Val: 9, Next: r2}
	res := AddTwoNumbers(l1, r1)

	assert.Equal(t, 1, res.Val, "Not same")
	assert.Equal(t, 0, res.Next.Val, "Not same")
	assert.Equal(t, 8, res.Next.Next.Val, "Not same")

	l2 = &ListNode{Val: 8}
	l1 = &ListNode{Val: 1, Next: l2}

	r1 = &ListNode{Val: 0}
	res = AddTwoNumbers(l1, r1)

	assert.Equal(t, 1, res.Val, "Not same")
	assert.Equal(t, 8, res.Next.Val, "Not same")
}

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := GroupAnagrams(input)
	fmt.Printf("%+v\n", result)
}

func TestValidSquare(t *testing.T) {
	p1 := []int{3, 0}
	p2 := []int{6, 0}
	p3 := []int{6, 3}
	p4 := []int{3, 3}
	result := validSquare(p1, p2, p3, p4)
	assert.True(t, result, "Not true")
}
