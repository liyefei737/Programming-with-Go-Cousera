package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Enter a string\n")
	var inputStr string
	fmt.Scanf("%s", &inputStr)
	// assuming ascii characters only a character takes 1 byte
	strLength := len(inputStr)
	inputStr = strings.ToLower(inputStr)
	if strLength > 3 &&
		string(inputStr[0]) == "i" && string(inputStr[strLength-1]) == "n" &&
		strings.Contains(inputStr, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
