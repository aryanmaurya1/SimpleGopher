package main

import "fmt"

func swap(arr []int, i, j int){
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
}

func pivot(arr []int, si, ei int) int {
	var p = ei // selecting as pivot
	var i int = si-1 
	var j int 
	for j = si; j<= p; j++ {
		if arr[j] <= arr[p] {
			i++
			swap(arr, i, j)
		}
	}
	return i // correct pivot
}

func quicksort(arr []int, si , ei int){
	if si >= ei{
		return 
	}
	var p = pivot(arr, si, ei)
	quicksort(arr, p+1, ei)
	quicksort(arr, si, p-1)
}

func main(){

	var arr = []int{7,8,9,4,5,6,1,2,3}
	var arr2 = []int{9,5,1,3,7,4,6,2,8}
	quicksort(arr, 0, len(arr)-1)
	quicksort(arr2, 4, 7)
	fmt.Println(arr)
	fmt.Println(arr2)

}