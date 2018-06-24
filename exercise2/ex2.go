package main

import (
	"fmt"
)

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func findMaxlen(names []string) int {
	max := 0
	for i := 0; i < len(names); i++ {
		if len(names[i]) > max {
			max = len(names[i])
		}
	}
	return max
}

func main() {
	maxLen := findMaxlen(names)
	result := make([][]string, maxLen)
	for _, name := range names {
		result[len(name)-1] = append(result[len(name)-1], name)
	}
	fmt.Printf("%v", result)
}
