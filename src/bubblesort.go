package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IntToOrdinal(i int) string {
	set := map[int]bool{1: true, 2: true, 3: true}
	suffix := map[int]string{1: "st", 2: "nd", 3: "rd"}
	lastDigit := i % 10
	if set[lastDigit] {
		return strconv.Itoa(i) + suffix[lastDigit]
	} else {
		return string(strconv.Itoa(i)) + "th"
	}
}

func main() {

	var inputInts []int

	scanner := bufio.NewScanner(os.Stdin)

	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter %s integer or press enter to exit\n", IntToOrdinal(i))
		scanner.Scan()
		userInput := scanner.Text()

		// exits when user presses enter
		if userInput == "" {
			break
		}

		inputInt, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		inputInts = append(inputInts, inputInt)
	}

	// given a list of integers, sort them and print the sorted list

	BubbleSort(inputInts)

	fmt.Printf("%v\n", inputInts)

}

func BubbleSort(slice []int) {

	//swap s[j] with s[j+1]
	swap := func(s []int, j int) { s[j+1], s[j] = s[j], s[j+1] }

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}

	}
}
