package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	animals := map[string]Animal{
		"cow": {
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
		"bird": {
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
		"snake": {
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		}}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		commands := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		switch commands[1] {
		case "eat":
			animals[commands[0]].eat()
		case "move":
			animals[commands[0]].move()
		case "speak":
			animals[commands[0]].speak()
		}
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) eat() {
	fmt.Println(animal.food)
}

func (animal Animal) move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) speak() {
	fmt.Println(animal.noise)
}
