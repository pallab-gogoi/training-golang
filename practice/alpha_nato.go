package main

import "fmt"

func main() {
	var word string
	fmt.Println("Enter the string: ")
	fmt.Scan(&word)
	covert(word)
}
func covert(word string) {
	for i := 0; i < len(word); i++ {
		x := string(word[i])
		switch x {
		case "a":
			fmt.Println("Alpha")
		case "b":
			fmt.Println("Bravo")
		case "c":
			fmt.Println("Charlie")
		case "d":
			fmt.Println("Delta")
		case "e":
			fmt.Println("Echo")
		case "f":
			fmt.Println("Foxtrot")
		case "g":
			fmt.Println("Golf")
		case "h":
			fmt.Println("Hotel")
		case "i":
			fmt.Println("India")
		case "j":
			fmt.Println("Juliet")
		case "k":
			fmt.Println("Kilo")
		case "l":
			fmt.Println("Lima")
		case "m":
			fmt.Println("Mike")
		case "n":
			fmt.Println("November")
		case "o":
			fmt.Println("Oscar")
		case "p":
			fmt.Println("Papa")
		case "q":
			fmt.Println("Quebec")
		case "r":
			fmt.Println("Romeo")
		case "s":
			fmt.Println("Siera")
		case "t":
			fmt.Println("Tango")
		case "u":
			fmt.Println("Uniform")
		case "v":
			fmt.Println("Victor")
		case "w":
			fmt.Println("Whiskey")
		case "x":
			fmt.Println("X-ray")
		case "y":
			fmt.Println("Yankee")
		case "z":
			fmt.Println("Zulu")

			//default:
			//fmt.Println("Invalid")
		}
	}
}
