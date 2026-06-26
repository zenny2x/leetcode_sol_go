package main

func getConcatenation(nums []int) []int {
	var nums2 []int

	nums2 = append(nums2, nums...)
	nums2 = append(nums2, nums...)

	return nums2
}
