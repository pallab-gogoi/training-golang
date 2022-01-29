/*
# Hamming

Welcome to Hamming on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Calculate the Hamming Distance between two DNA strands.

Your body is made up of cells that contain DNA. Those cells regularly wear out and need replacing,
which they achieve by dividing into daughter cells. In fact, the average human body experiences
about 10 quadrillion cell divisions in a lifetime!

When cells divide, their DNA replicates too. Sometimes during this process mistakes happen and
single pieces of DNA get encoded with the incorrect information. If we compare two strands of DNA and count
the differences between them we can see how many mistakes occurred. This is known as the "Hamming Distance".

We read DNA using the letters C,A,G and T. Two strands might look like this:

    GAGCCTACTAACGGGAT
    CATCGTAATGACGGCCT
    ^ ^ ^  ^ ^    ^^

They have 7 differences, and therefore the Hamming Distance is 7.

The Hamming Distance is useful for lots of things in science, not just biology, so it's a nice phrase to be familiar with :)

## Implementation notes

The Hamming distance is only defined for sequences of equal length, so
an attempt to calculate it between sequences of different lengths should
not work.

You may be wondering about the `cases_test.go` file. We explain it in the
[leap exercise][leap-exercise-readme].

[leap-exercise-readme]: https://github.com/exercism/go/blob/main/exercises/practice/leap/.docs/instructions.md

## Source

### Created by

- @levicook

### Contributed to by

- @alebaffa
- @bitfield
- @da-edra
- @ekingery
- @ferhatelmas
- @hilary
- @hpurmann
- @ilmanzo
- @kytrinyx
- @leenipper
- @mchoube
- @paul-nelson-baker
- @petertseng
- @robphoenix
- @sebito91
- @soniakeys
- @tleen
- @tompao

### Based on

The Calculating Point Mutations problem at Rosalind - http://rosalind.info/problems/hamm/
*/

package main

import "fmt"

func main() {
	strand1 := "GAGCCTACTAACGGGAT"
	strand2 := "CATCGTAATGACGGCCT"
	fmt.Println(checkDistance(strand1, strand2))

}
func checkDistance(s1, s2 string) int {
	hd := 0
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				hd++
			}

		}
		return hd
	}
	return 0
}
