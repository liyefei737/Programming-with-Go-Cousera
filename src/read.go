package main

import (
	"fmt"
  "os"
  "bufio"
  "strings"
)

func main() {

  var fileName string

  type Name struct {
    fname string
    lname string
  }

  var names []Name

  stdinScanner := bufio.NewScanner(os.Stdin)

  fmt.Println("Enter the path(name) of the file you want to read")
  stdinScanner.Scan()
  fileName = stdinScanner.Text()

  fd, err := os.Open(fileName)

  if err != nil {
    fmt.Println(err)
    return
  }

  fileScanner := bufio.NewScanner(fd)

  for fileScanner.Scan() {
    line := fileScanner.Text()  // get a line from the file
    words := strings.Split(line, " ")
    n := Name{fname: words[0], lname: words[1]}
    names = append(names, n)
  }

  // if error happens during Scan, fileScanner.Err() will return non-nil value
  err = fileScanner.Err()

  if err != nil {
    fmt.Println("error reading the file")
    return
  }

  fmt.Printf("%v\n", names)

  fd.Close()

}
