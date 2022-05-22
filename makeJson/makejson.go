package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// receive user input
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please entre a name: ")
	name, err := inputReader.ReadString('\n')
	if err == nil {
		name = strings.TrimSpace(name)
	} else {
		fmt.Println("Something is wrong, Please start the program again!")
		return
	}
	fmt.Println("Please entre an address: ")
	address, err := inputReader.ReadString('\n')
	if err == nil {
		address = strings.TrimSpace(address)
	} else {
		fmt.Println("Something is wrong, Please start the program again!")
		return
	}

	// create a map for user input
	userInfo := map[string]string{
		"name":    name,
		"address": address,
	}

	// JSON object from the map
	userInfoJson, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("Something is wrong, Please start the program again!")
		return
	}
	// Print out the json object
	fmt.Printf("user input in JSON format -> %v\n", string(userInfoJson))

}
