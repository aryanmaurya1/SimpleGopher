package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	acc, initV, initD := getUserVariables()
	displacement := GenDisplaceFn(acc, initV, initD)

	var userTime string
	fmt.Printf("Calculating displacement using %.2f acceleration, %.2f initial velocity and %.2f initial displacement...", acc, initV, initD)
	for {
		userTime = getInput("\nEnter time: ")
		if strings.ToLower(userTime) == "x" {
			fmt.Println("Closing program.")
			os.Exit(1)
		}
		convUserTime := parseStringIntoFloat(userTime)
		fmt.Printf("\nAfter %.2fs, displacement is: %.2f", convUserTime, displacement(convUserTime))
	}
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5*a*math.Pow(t, 2) + v0*t + s0)
	}
	return fn
}

func getUserVariables() (float64, float64, float64) {
	acc := parseStringIntoFloat(getInput("Enter acceleration: "))
	initV := parseStringIntoFloat(getInput("Enter initial velocity: "))
	initD := parseStringIntoFloat(getInput("Enter initial displacement: "))

	return acc, initV, initD
}

func parseStringIntoFloat(str string) float64 {
	convFloat, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return convFloat
}

func getInput(instr string) string {
	fmt.Print(instr)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.TrimRight(input, "\r\n")
}