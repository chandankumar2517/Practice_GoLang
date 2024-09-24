package maximumsubarray

func GetMaximumSubArray(nums []int) (int, []int) {

	// {-2,1,-3,4,-1,2,1,-5,4}

	max_sum := nums[0]
	max_global := nums[0]

	start, end := 0, 0 // To track the start and end of the subarray
	tempStart := 0     // To track the potential start of a new subarray

	for i := 1; i < len(nums); i++ {

		// Either include the current element in the current subarray or start a new subarray
		if nums[i] > max_sum+nums[i] {
			max_sum = nums[i]
			tempStart = i // Start a new subarray from index i
		} else {
			max_sum += nums[i]
		}

		if max_sum > max_global {
			max_global = max_sum

			start = tempStart // Update the start of the max subarray
			end = i           // Update the end of the max subarray
		}
	}

	return max_sum, nums[start : end+1]
}
