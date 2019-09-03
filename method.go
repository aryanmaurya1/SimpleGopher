package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	Food       string
	Locomotion string
	Noise      string
}

func setData(cow, bird, snake *animal) {
	cow.Food = "grass"
	cow.Locomotion = "walk"
	cow.Noise = "moo"

	bird.Food = "worms"
	bird.Locomotion = "fly"
	bird.Noise = "peep"

	snake.Food = "mice"
	snake.Locomotion = "slither"
	snake.Noise = "hsss"
}

func (a animal) Eat() string {
	return a.Food
}
func (a animal) Move() string {
	return a.Locomotion
}
func (a animal) Speak() string {
	return a.Noise
}

func main() {
	var cow, bird, snake animal
	setData(&cow, &bird, &snake)

	var pet, method string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(">> ")
	for scanner.Scan() {
		data := strings.Split(string(scanner.Text()), " ")
		pet = data[0]
		method = data[1]
		switch pet {
		case "cow":
			{
				switch method {
				case "Eat":
					fmt.Println(cow.Eat())
				case "Move":
					fmt.Println(cow.Move())
				case "Speak":
					fmt.Println(cow.Speak())
				}
			}
		case "bird":
			{
				switch method {
				case "Eat":
					fmt.Println(bird.Eat())
				case "Move":
					fmt.Println(bird.Move())
				case "Speak":
					fmt.Println(bird.Speak())
				}
			}
		case "snake":
			{
				switch method {
				case "Eat":
					fmt.Println(snake.Eat())
				case "Move":
					fmt.Println(snake.Move())
				case "Speak":
					fmt.Println(snake.Speak())
				}
			}
		}
		fmt.Printf(">> ")
	}
}
