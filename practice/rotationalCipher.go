package main

import "fmt"

func main() {
	fmt.Println(RotationCipher("OMG"))
}

var alphabet = map[int]string{ // defined a map for alphabets
	1: "a", "A",
	2:  "b",
	3:  "c",
	4:  "d",
	5:  "e",
	6:  "f",
	7:  "g",
	8:  "h",
	9:  "i",
	10: "j",
	11: "k",
	12: "l",
	13: "m",
	14: "n",
	15: "o",
	16: "p",
	17: "q",
	18: "r",
	19: "s",
	20: "t",
	21: "u",
	22: "v",
	23: "w",
	24: "x",
	25: "y",
	26: "z",
}

func GetKey(letter string) (key int) {
	for index, value := range alphabet {
		if value == letter {
			key = index
			break
		}
	}
	return
}
func RotationCipher(original string) (encrypted string) {
	for _, value := range original {
		shiftedKey := GetKey(string(value)) + 5 //ROT5
		if shiftedKey > 26 {
			shiftedKey = shiftedKey % 26
		}
		encrypted += alphabet[shiftedKey]
	}
	return
}
