package main

import "fmt"

func printMap(theMap map[string]string) {
	for key, val := range theMap {
		fmt.Println(key, " -> ", val)
	}
}

func main() {

	// var colors map[string]string		// First way to declare a map

	// colors := make(map[string]string)	// Another way to declare a map

	colors := map[string]string{
		"Red":    "ff0000",
		"Yellow": "f763fg",
		"Black":  "000000",
		"White":  "ffffff",
	}

	colors["Pink"] = "f56dfy"

	delete(colors, "Red")

	// fmt.Println(colors)
	printMap(colors)
}
