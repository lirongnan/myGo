package main

import "fmt"

func main() {
	var nums, target = []int{2, 7, 11, 15}, 9
	var res = twoSum(nums, target)
	fmt.Print(res)
}

func twoSum(nums []int, target int) []int {
	//
	var res = []int{0, 0}
	for i := 0; i < len(nums)-1; i++ {
		e := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == e {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}
