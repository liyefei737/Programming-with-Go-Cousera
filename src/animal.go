package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Animal methods
func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	animalGroup := map[string]Animal{
		"cow":   Animal{food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  Animal{food: "worms", locomotion: "fly", noise: "peep"},
		"snake": Animal{food: "mice", locomotion: "slither", noise: "hsss"},
	}

	stdinScanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("> ")
		stdinScanner.Scan()
		line := stdinScanner.Text()

		if line == "" {
			fmt.Println("exitting...")
			break
		}

		words := strings.Split(line, " ")
		if len(words) != 2 {
			fmt.Println("request takes 2 words exactly separated by a single space")
			continue
		}

		selectetAnimal, ok := animalGroup[words[0]]

		if !ok {
			fmt.Println("unrecognizable animal entered")
			continue
		}

		switch words[1] {
		case "Eat":
			selectetAnimal.Eat()
		case "Move":
			selectetAnimal.Move()
		case "Speak":
			selectetAnimal.Speak()
		default:
			fmt.Println("unrecognizable animal action")
			continue

		}

	}

}
