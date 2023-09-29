package main

func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	direction := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if direction == 0 {
				direction = 1
			} else if direction == -1 {
				return false
			}
		} else if nums[i] < nums[i-1] {
			if direction == 0 {
				direction = -1
			} else if direction == 1 {
				return false
			}
		}
	}
	return true
}
