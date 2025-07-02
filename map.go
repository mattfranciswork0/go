package main

import "fmt"

//map  are passed by  value (hence  yuou dont need a pointer in
// the arguemnt of printMap, whereas structs are pass by reference like
func runMapExample() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["yellow"] = "laksjdf"
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
