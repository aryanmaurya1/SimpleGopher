package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {

	return func(t float64) float64 {
		return (1.0/2)*(a*t*t) + (v * t) + s
	}
}
func main() {
	var a, v, s float64
	fmt.Println("Enter Values of Acceleration, Velocity, Initial Displacement : ")
	fmt.Scanf("%f%f%f", &a, &v, &s)
	f := GenDisplaceFn(a, v, s)
	var t float64
	fmt.Println("Enter Time : ")
	fmt.Scanf("%f", &t)
	fmt.Printf("Displacement after %0.2f unit time is %0.2f units.\n", t, f(t))
}
