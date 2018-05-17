package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	printColours(colours)
}

func printColours(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is:", hex)
	}
}
