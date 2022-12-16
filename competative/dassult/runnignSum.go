package main

func main() {
	// x := []int{-4, 3, 2, 1}
	// x := []int{-1, -2, -5}
	// x := []int{3, -6, 5, -2, 1}
	print("got1 ", minRunningSum(x, 0))
	print("got2 ", getMinX(x))
	print("got3 ", minStartValue(x))
}

func minRunningSum(nums []int, sum int) int {
	orgSum := sum
	for _, v := range nums {
		sum += v
		if sum < 1 {
			orgSum += Abs(v) + 1
			break
		}
	}
	if !(sum < 1) {
		if nums[0] > 0 && (orgSum-nums[0]) > 0 {
			return orgSum - nums[0]
		}
		return orgSum
	} else {
		return minRunningSum(nums, orgSum)
	}
}

func Abs(v int) int {
	if v > 0 {
		return v
	} else {
		return v * -1
	}
}

func getMinX(arr []int) int {
	x := 1
	for _, v := range arr {
		x = x - v
		if x < 1 {
			x = 1
		}
	}
	return x
}

func minStartValue(nums []int) int {
	prefixSum, result := 0, 1 //int(0), int(1)
	for _, num := range nums {
		prefixSum += num
		result = max(result, 1-prefixSum)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
