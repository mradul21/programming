package main

// Given a collection of distinct integers, return all possible permutations.

import "fmt"

func permute(nums []int) [][]int {
	var out [][]int

	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}

	for ii := 0; ii < len(nums); ii++ {
		nums1 := make([]int, len(nums))
		copy(nums1, nums)
		tmp1 := append(nums1[:ii], nums1[ii+1:]...)
		temp := permute(tmp1)
		var out1 []int
		for _, out1 = range temp {
			out1 = append([]int{nums[ii]}, out1...)
			out = append(out, out1)
		}

	}

	return out
}

func main() {
	a := []int{1, 2, 3, 4}
	b := permute(a)

	fmt.Println("%v", b)
}
