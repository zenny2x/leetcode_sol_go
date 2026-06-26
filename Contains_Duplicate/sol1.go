package main

func hasDuplicate(nums []int) bool {
	for i, v := range nums {
		for i2, v2 := range nums {
			if i != i2 && v == v2 {
				return true
			}
		}
	}
	return false
}
