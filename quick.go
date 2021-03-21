package main

import (
	"fmt"
)

func main() {
	fmt.Println("Quick Sort Start!")

	// 設問用
	list := []int{4, 2, 5, 6, 1, 9, 10, 7}

	// ソート
	list = quicksort(list)

	// 表
	fmt.Println("result:")
	fmt.Println(list)
}

func quicksort(seq []int) []int {
	if len(seq) < 1 {
		return seq
	}
	pivot := seq[0]
	left := []int{}
	right := []int{}
	for i := 1; i < len(seq); i++ {
		if seq[i] <= pivot {
			left = append(left, seq[i])
		} else {
			right = append(right, seq[i])
		}
	}
	left = quicksort(left)
	right = quicksort(right)
	left = append(left, pivot)
	left = append(left, right...)
	return left
}
