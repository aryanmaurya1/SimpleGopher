package main

import "fmt"

func swap(input []int, index, nextIndex int) {
	temp := input[index]
	input[index] = input[nextIndex]
	input[nextIndex] = temp
}
func bubbleSort(input []int) {
	length := len(input)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if input[j] > input[j+1] {
				swap(input, j, j+1)
			}
		}
	}
}

func main() {
	var input []int
	fmt.Println("Enter 10 Integers : ")
	for i := 0; i < 10; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		input = append(input, temp)
	}
	fmt.Println("Original Slice : ", input)
	bubbleSort(input)
	fmt.Println("Sorted Slice : ", input)

}
