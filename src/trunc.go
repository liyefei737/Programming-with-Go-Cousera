package main

import (
  "fmt"
  "strconv"
)

func main() {
   fmt.Printf("Enter a floating point number\n")
   var floatNum float64
	 fmt.Scanf("%f", &floatNum)
   truncatedFloat := strconv.Itoa(int(floatNum))
   fmt.Printf("the truncated version is " + truncatedFloat + "\n")
}
