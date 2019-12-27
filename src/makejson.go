package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var name string
	var addr string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter an name")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("Enter an address")
	scanner.Scan()
	addr = scanner.Text()

	var userInfo map[string]string
	userInfo = make(map[string]string)

	userInfo["name"] = name
	userInfo["address"] = addr
	jsonUserInfo, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("json Marshaling failed")
		return
	}

	fmt.Printf("%s\n", jsonUserInfo)

}
