package main

import "fmt"

var outputArray [][]int

func subsets(arr []int) int {
	if len(arr) == 0 {
		outputArray = append(outputArray, []int{})
		return len(outputArray)
	}
	NumberSubsets := subsets(arr[1:])
	for i := 0; i < NumberSubsets; i++ {
		var temp = make([]int, len(outputArray[i]), len(outputArray[i]))
		copy(temp, outputArray[i])
		temp = append(temp, arr[0])
		outputArray = append(outputArray, temp)
	}
	return len(outputArray)
}

func main() {
	var input = []int{7, 4, 1, 2, 3, 6, 5, 4, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	subsets(input)
	fmt.Println(outputArray)
}
