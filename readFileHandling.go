package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type person struct{
	firstName string
	lastName string
}

func main() {
	var arr []person
	var filename string
	var temp person

	fmt.Printf("Enter Filename : ")
	fmt.Scanf("%s",&filename)

	data, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error in opening file.")
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		cName := strings.Split(string(line)," ")
		temp.firstName = cName[0]
		temp.lastName = cName[1]
		arr = append(arr, temp)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i].firstName, arr[i].lastName)
	}
}
