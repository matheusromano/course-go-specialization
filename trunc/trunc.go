package main

import (
	"fmt"
)

func main() {
	fmt.Println("Digit the number: ")
	var usernum float64
	fmt.Scan(&usernum)
	fmt.Printf("Trucated number: %d", int(usernum))

}
