package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Speak()
	Eat()
	Move()
	GetName() string
}

type Cow struct {
	name       string
	food       string
	locomotion string
	sound      string
}

type Snake struct {
	name       string
	food       string
	locomotion string
	sound      string
}

type Bird struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (a Cow) Speak() {
	fmt.Println(a.sound)
}
func (a Snake) Speak() {
	fmt.Println(a.sound)
}
func (a Bird) Speak() {
	fmt.Println(a.sound)
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}
func (a Snake) Eat() {
	fmt.Println(a.food)
}
func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) GetName() string {
	return a.name
}
func (a Snake) GetName() string {
	return a.name
}
func (a Bird) GetName() string {
	return a.name
}

var data []Animal

func newAnimal(cmd1, cmd2 string) {
	var temp Animal
	switch cmd2 {
	case "Cow", "cow":
		var t = new(Cow)
		t = &Cow{name: cmd1, food: "grass", locomotion: "walk", sound: "moo"}
		temp = t
	case "Snake", "snake":
		var t = new(Snake)
		t = &Snake{name: cmd1, food: "mice", locomotion: "slither", sound: "hsss"}
		temp = t

	case "Bird", "bird":
		var t = new(Bird)
		t = &Bird{name: cmd1, food: "worms", locomotion: "fly", sound: "peep"}
		temp = t
	}
	data = append(data, temp)
}

func query(cmd1, cmd2 string) {
	var foundTag bool = false
	for i := 0; i < len(data); i++ {
		if data[i].GetName() == cmd1 {
			switch cmd2 {
			case "food", "Food", "eat", "Eat":
				data[i].Eat()
				foundTag = true
			case "sound", "Sound", "speak", "Speak":
				data[i].Speak()
				foundTag = true
			case "locomotion", "Locomotion", "move", "Move":
				data[i].Move()
				foundTag = true
			default:
				fmt.Println("Error !!!")
			}
		}
	}
	if !foundTag {
		fmt.Println("Animal Not Found !!!")
	}
}

func transfer(mode, cmd1, cmd2 string) {
	if mode == "newanimal" {
		newAnimal(cmd1, cmd2)

	} else if mode == "query" {
		query(cmd1, cmd2)
	}
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Printf(">> ")
	for scanner.Scan() {
		line := strings.Split(string(scanner.Text()), " ")
		if len(line) != 3 {
			fmt.Println("Recheck Command !!")
			fmt.Printf(">> ")
			continue
		}
		mode := line[0]
		cmd2 := line[1]
		cmd3 := line[2]
		transfer(mode, cmd2, cmd3)
		fmt.Printf(">> ")
	}
}
