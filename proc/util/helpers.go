package util

// ReverseBoolArray reverses passed bool array
func ReverseBoolArray(arr []bool) []bool {
	for i := 0; i < len(arr)-1; i++ {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
	return arr
}

// ArrayBoolToInt makes array of int from array of bool
// [false, true, true] -> [0, 1, 1]
func ArrayBoolToInt(arr []bool) []int {
	res := make([]int, len(arr))
	for i, value := range arr {
		if value {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}
