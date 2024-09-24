package twosum

func TwoSum(nums []int, target int) [][]int {

	components := make(map[int][]int)

	result := [][]int{}

	for i, num := range nums {

		if componentIndices, exists := components[target-num]; exists {

			for _, componentIndex := range componentIndices {

				result = append(result, []int{componentIndex, i})
			}

		}

		components[num] = append(components[num], i)
		/*for _, item := range components[num] {
		    fmt.Println("item  ", item)
		}*/

	}

	return result
}
