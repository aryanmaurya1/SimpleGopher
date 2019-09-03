package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func getArray() []int {
	fmt.Printf("Enter the numbers : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Split(string(scanner.Text()), " ")
	var intArr = make([]int, len(arr), len(arr))
	for index, val := range arr {
		temp, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("Error, recheck input !!! ")
			os.Exit(0)
		}
		intArr[index] = temp
	}
	return intArr
}

func sortArr(arr []int, wg *sync.WaitGroup) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
	wg.Done()
}

func mergeArr(arr []int, arr2 []int) []int {
	var lenMax = len(arr) + len(arr2)
	mergedArr := make([]int, lenMax, lenMax)
	var i, j, k int
	for i < len(arr) && j < len(arr2) {
		if arr[i] > arr2[j] {
			mergedArr[k] = arr2[j]
			j++
		} else {
			mergedArr[k] = arr[i]
			i++
		}
		k++
	}
	for i < len(arr) {
		mergedArr[k] = arr[i]
		i++
		k++
	}
	for j < len(arr2) {
		mergedArr[k] = arr2[j]
		j++
		k++
	}
	return mergedArr
}

func main() {
	arr := getArray()
	arr1 := arr[0 : len(arr)/4]
	arr2 := arr[len(arr)/4 : len(arr)/2]
	arr3 := arr[len(arr)/2 : (len(arr)*3)/4]
	arr4 := arr[(len(arr)*3)/4:]

	wg.Add(4)
	go sortArr(arr1, &wg)
	go sortArr(arr2, &wg)
	go sortArr(arr3, &wg)
	go sortArr(arr4, &wg)
	wg.Wait()

	merged1 := mergeArr(arr1, arr2)
	merged2 := mergeArr(arr3, arr4)
	fmt.Println(mergeArr(merged1, merged2))

}
