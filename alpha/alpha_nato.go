package main

import (
	"time"
)

func main() {
	//var word string
	//return("Enter the string: ")
	//fmt.Scan(&word)
	//covert(word)
}
func covert(word string) string {
	for i := 0; i < len(word); i++ {
		//fmt.Print()("  ")
		x := string(word[i])
		time.Sleep(time.Second)
		switch x {
		case "a", "A":
			return ("Alpha")
		case "b", "B":
			return ("Bravo")
		case "c", "C":
			return ("Charlie")
		case "d", "D":
			return ("Delta")
		case "e", "E":
			return ("Echo")
		case "f", "F":
			return ("Foxtrot")
		case "g", "G":
			return ("Golf")
		case "h", "H":
			return ("Hotel")
		case "i", "I":
			return ("India")
		case "j", "J":
			return ("Juliet")
		case "k", "K":
			return ("Kilo")
		case "l", "L":
			return ("Lima")
		case "m", "M":
			return ("Mike")
		case "n", "N":
			return ("November")
		case "o", "O":
			return ("Oscar")
		case "p", "P":
			return ("Papa")
		case "q", "Q":
			return ("Quebec")
		case "r", "R":
			return ("Romeo")
		case "s", "S":
			return ("Siera")
		case "t", "T":
			return ("Tango")
		case "u", "U":
			return ("Uniform")
		case "v", "V":
			return ("Victor")
		case "w", "W":
			return ("Whiskey")
		case "x", "X":
			return ("X-ray")
		case "y", "Y":
			return ("Yankee")
		case "z", "Z":
			return ("Zulu")

			//default:
			//return("Invalid")
		}

	}
	return "c"
}
