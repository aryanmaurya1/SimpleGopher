package main

import "fmt"
import "strconv"
import "sort"

func main() {
	var arr []int = make([]int, 0, 3)
	var input string
	for {
		fmt.Scanf("%s", &input)
		if input == "X" {
			break
		}
		temp, _ := strconv.Atoi(input)
		arr = append(arr, temp)
	}
	sort.Ints(arr)
	fmt.Println(arr)
}
