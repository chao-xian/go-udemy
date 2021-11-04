package main

import "fmt"

func main() {
	// literal map[key-type]value-type
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#408234",
	}

	// initialised no-value
	// var colours map[string]string

	// similiar to initialising
	// colours := make(map[string]string)

	colours["white"] = "#ffffff"
	fmt.Println(colours["white"])

	delete(colours, "white")

	fmt.Println(colours)
}
