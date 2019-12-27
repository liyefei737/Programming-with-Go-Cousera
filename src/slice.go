package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var inputStr string
	sli := make([]int, 0, 3) // create a slice of length 3

	fmt.Printf("Enter an integer\n")
	fmt.Scanf("%s", &inputStr)

	for inputStr != "x" {
		i, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("a value that is not an integer nor x is entered. Please enter a valid integer.")
		} else {
			sli = append(sli, i)
			sort.Ints(sli)
			fmt.Println(sli)
		}

		fmt.Printf("Enter an integer\n")
		fmt.Scanf("%s", &inputStr)
	}
}
