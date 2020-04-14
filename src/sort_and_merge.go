package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Enter integers separated by space: \n")

		scanner.Scan()
		inputText := strings.TrimSpace(scanner.Text())

		if inputText == "" {
			continue
		}

		inputTokens := strings.Split(inputText, " ")

		if len(inputTokens) < 4 {
			fmt.Printf("Enter again. Less than 4 integers supplied.\n")
			continue
		}

		var intList []int
		for _, s := range inputTokens {
			number, _ := strconv.Atoi(s)
			intList = append(intList, number)
		}

		// rounding down e.g. len = 12, partionSize is 3
		// if len is not a mutiple of 4, the last segment has more elments than the earlier ones
		partitionSize := len(intList) / 4

		sortedChan := make(chan []int)

		go sortSubArr(intList[:partitionSize], sortedChan)
		go sortSubArr(intList[partitionSize:2*partitionSize], sortedChan)
		go sortSubArr(intList[2*partitionSize:3*partitionSize], sortedChan)
		go sortSubArr(intList[3*partitionSize:], sortedChan)

		sorted0 := <-sortedChan
		sorted1 := <-sortedChan
		sorted2 := <-sortedChan
		sorted3 := <-sortedChan

		merged := mergeSortedSl(sorted0, sorted1, sorted2, sorted3)
		fmt.Printf("overal merged: %v.\n", merged)
	}

}

func sortSubArr(sl []int, outputChan chan []int) {
	fmt.Printf("sorting slice in a goroutine: %v.\n", sl)
	sort.Ints(sl)
	outputChan <- sl
}

func mergeSortedSl(slices ...[]int) []int {
	if len(slices) == 0 {
		return nil
	}
	if len(slices) == 1 {
		return slices[0]
	}

	mergedSl := slices[0]

	for i := 1; i < len(slices); i++ {
		mergedSl = merge2SortedSl(mergedSl, slices[i])

	}

	return mergedSl
}

func merge2SortedSl(sl1, sl2 []int) []int {
	len1, len2 := len(sl1), len(sl2)
	mergedSl := []int{}

	i, j := 0, 0 //index for looping through sl1 and sl2
	for i != len1 && j != len2 {
		if sl1[i] < sl2[j] {
			mergedSl = append(mergedSl, sl1[i])
			i++
		} else {
			mergedSl = append(mergedSl, sl2[j])
			j++
		}
	}

	// one of the slice exausts first, slice with smallest upper bound value
	// we want to find the slice that did not exaust and add its values left to the end
	// these for loops are like if statement
	for i < len1 {
		mergedSl = append(mergedSl, sl1[i])
		i++
	}

	for j < len2 {
		mergedSl = append(mergedSl, sl2[j])
		j++
	}

	return mergedSl

}
