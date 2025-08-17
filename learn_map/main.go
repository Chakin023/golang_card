package main

import "fmt"

func main() {
	// 1create empty map
	// var colors map[string]string
	// colors := make(map[string]string)
	//
	// 1.1map with values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// 2add key, value to map
	// colors["white"] = "#ffffff"

	// 3delete key, value from map
	// delete(colors, "white")

	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Hex code for", k, "is", v)
	}
}
