package main

func maxArea(height []int) int {
	maxarea := 0
	l := 0
	r := len(height) - 1

	for l < r {
		maxarea = max(maxarea, min(height[l], height[r])*(r-l))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return maxarea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
