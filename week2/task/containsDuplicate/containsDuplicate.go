package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	// megurutkan data
	sort.Ints(nums)

	// mengecek apakah
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums1)) //true

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2)) //false

	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums3)) //true

}
