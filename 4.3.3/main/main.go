package main

import (
	"fmt"
	"unicode/utf8"
)

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}


func main() {

	var max int

	for _, value := range names {
		if len := utf8.RuneCountInString(value); len > max {
			max = len
		}
	}

	counter := make([][]string, max)

	for _, value := range names {
		len := utf8.RuneCountInString(value)
		counter[len-1]=append(counter[len-1], value)
	}

	fmt.Printf("%q", counter)

}
