package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
	// outputs change to "Bye" because slices are pointers to arrays really
	// and the pointer gets copied but points to the same array

	// Slices are called reference types, so are maps, channels, pointers and functions
	// Value types are int, float, string, bool and structs
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
