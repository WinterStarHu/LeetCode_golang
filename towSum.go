package main

import "fmt"

func main() {
	var nums = []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
func twoSum(nums []int, target int) []int {
	dMap := make(map[int]int)
	for idx, val := range nums {
		minus := target - val
		if _, ok := dMap[minus]; ok {
			return []int{dMap[minus], idx}
		} else {
			dMap[val] = idx
		}
	}
	return nil
}
