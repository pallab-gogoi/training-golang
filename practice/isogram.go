/*
# Isogram

Welcome to Isogram on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Determine if a word or phrase is an isogram.

An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

Examples of isograms:

- lumberjacks
- background
- downstream
- six-year-old

The word *isograms*, however, is not an isogram, because the s repeats.

## Source

### Created by

- @erizocosmico

### Contributed to by

- @alebaffa
- @bitfield
- @da-edra
- @ekingery
- @ferhatelmas
- @hilary
- @ilmanzo
- @leenipper
- @petertseng
- @robphoenix
- @sebito91
- @strangeman
- @tleen

### Based on

Wikipedia - https://en.wikipedia.org/wiki/Isogram
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var word string
	fmt.Println("Enter the word: ")
	fmt.Scan(&word)
	if isIsogram(word) == true {
		fmt.Println("The given string is a Isogram.")
	} else {
		fmt.Println("The given string is NOT  a Isogram")
	}
}
func isIsogram(word string) bool {
	var sample = ""
	for _, char := range word {
		if string(char) == " " || string(char) == "-" {
			continue
		}
		if strings.Contains(sample, string(unicode.ToUpper(char))) || strings.Contains(sample, string(unicode.ToLower(char))) {
			//fmt.Printf("%c is already present", char)
			return false
		}
		sample += string(char)
	}
	return true
	//panic("Implement isIsogram function")

}
