package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name_struc struct {
	fname string
	lname string
}

func main() {
	var my_file string
	var my_slice []name_struc
	fmt.Scan(&my_file)
	file, err := os.Open(my_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var name name_struc
		string_split := strings.Split(scanner.Text(), " ")
		name.fname = string_split[0]
		name.lname = string_split[1]
		my_slice = append(my_slice, name)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, name := range my_slice {
		fmt.Println(name.fname, " ", name.lname)
	}
}
