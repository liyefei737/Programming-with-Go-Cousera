package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

//Cow methods
func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

//Bird methods
func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

//Snake methods
func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	animalGroup := map[string]Animal{}

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
		if len(words) != 3 {
			fmt.Println("command takes 3 words exactly separated by a single space")
			continue
		}

		command := words[0]
		animalName := words[1]

		switch command {
		case "newanimal":
			animalType := words[2]
			switch animalType {
			case "cow":
				animalGroup[animalName] = new(Cow)
			case "bird":
				animalGroup[animalName] = new(Bird)
			case "snake":
				animalGroup[animalName] = new(Snake)
			default:
				fmt.Println("unrecognizable animal type")
				continue
			}
      fmt.Println("Created it!")

		case "query":
			_, isExisting := animalGroup[animalName]
			if !isExisting {
				fmt.Println(animalName + animalName + "does not exist")
				continue
			}
			animalAction := words[2]
			switch animalAction {
			case "eat":
				animalGroup[animalName].Eat()
			case "move":
				animalGroup[animalName].Move()
			case "speak":
				animalGroup[animalName].Speak()
			default:
				fmt.Println("unrecognizable animal action")
				continue

			}
		default:
			fmt.Println("unrecognizable command")
			continue

		}

	}

}
