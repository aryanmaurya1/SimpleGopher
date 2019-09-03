package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the larrysArray function below.
func isEqual(arr, temp []int32) bool {
	for index, item := range temp {
		if arr[index] != item {
			return false
		}
	}
	return true
}

func sortThree(arr []int32) bool {
	temp := make([]int32, 3, 3)
	var firstTime = false
	copy(temp, arr)
	for {
		if arr[2] >= arr[1] && arr[2] >= arr[0] {
			return true
		}
		if isEqual(arr, temp) && firstTime {
			return false
		}
		k := arr[0]
		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = k
		firstTime = true
	}

}

func larrysArray(A []int32) string {
	var n = int32(len(A))
	var result string
	var i int32
	var j int32
	for i = 0; i < n-2; i++ {
		for j = 0; j < n-2-i; j++ {
			if i == n-3 {
				var isSorted = sortThree(A[0:4])
				if isSorted {
					result = "YES"
				} else {
					result = "NO"
				}
			} else {
				sortThree(A[j : j+3])
			}
		}
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		ATemp := strings.Split(readLine(reader), " ")

		var A []int32

		for i := 0; i < int(n); i++ {
			AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			A = append(A, AItem)
		}

		result := larrysArray(A)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
