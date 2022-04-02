package main

/*
This is going to show the example of Map
*/
import "fmt"

func main() {
	// TODO 3: create a map called colors with no value - example 2
	// colors := make(map[string]string)

	// Add the red with hex
	// colors["red"] = "#ff0000"

	// Delete the red color in map
	// delete(colors, "red")

	// TODO 2: create a map called colors with no value - example 1
	// var colors map[string]string

	// fmt.Println(colors)

	// TODO 1: simple map example
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	pringMap(colors)
}

func pringMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
