package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 17))
}

func twoSum(nums []int, target int) []int {

	for x := 0; x < len(nums); x++ {
		i := target - nums[x]
		for y := x + 1; y < len(nums); y++ {
			if i == nums[y] {
				return []int{x, y}
			}
		}
	}

	return []int{}
}
