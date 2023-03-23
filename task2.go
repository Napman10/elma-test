package main

import "fmt"

func rotation(A []int) []int {
	result := make([]int, 0, len(A))

	for i := range A {
		j := i - 1

		if j < 0 {
			j = len(A) - 1
		}

		result = append(result, A[j])
	}

	return result
}

func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	offset := K % len(A)

	for i := 0; i < offset; i++ {
		A = rotation(A)
	}

	return A
}

func SolutionVer2(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	offset := K % len(A)
	result := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		newIdx := i + offset

		if newIdx >= len(A) {
			newIdx -= len(A)
		}

		result[newIdx] = A[i]
	}

	return result
}

func ShowTask2() {
	arr := []int{4, 6, -8, -4, 0, 3}

	fmt.Printf("source - %v\n", arr)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d rotations - %v\n", i, Solution(arr, i))
	}
}
