package main

import "fmt"

func main() {
	var nums =[]int{3,2,4}
	target := 6
	fmt.Println(twoSum(nums,target))
}
func twoSum(nums []int, target int) []int {
	var dMap = make(map[int]int)
	for idx, val := range nums {
		_, ok := dMap[target-val]
		if ok {
			return []int{idx, dMap[target-val]}
		} else {
			dMap[val] = idx
		}
	}
	return nil
}

