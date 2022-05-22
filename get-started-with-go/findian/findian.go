package main

import (
	"fmt"
	"strings"
)

func main() {

	for {
		fmt.Print("Please, enter a string: ")

		var s string

		_, err := fmt.Scanf("%s", &s)

		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Printf("Input is %s\n", s)
			s1 := strings.ToLower(s)
			s2 := strings.TrimSpace(s1)
			if strings.HasPrefix(s2, "i") && strings.Contains(s2, "a") && s2[len(s2)-1:] == "n" {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}
		}

	}

}
