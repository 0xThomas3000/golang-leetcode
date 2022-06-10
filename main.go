package main

import "fmt"

func twoSums(nums []int, target int) []int {
	var arr []int // Declaring Slice
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr = append(arr, i, j)
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(twoSums([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSums([]int{3, 2, 4}, 6))
	fmt.Println(twoSums([]int{3, 3}, 6))
}
