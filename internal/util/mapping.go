package util

import "strconv"

func MapStringArrayToIntArray(array []string) ([]int, error) {
	nums := make([]int, len(array))
	for i, value := range array {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return nums, nil
}
