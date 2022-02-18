/*
Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with  places after the decimal.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to  are acceptable.

Example

There are  elements, two positive, two negative and one zero. Their ratios are ,  and . Results are printed as:

0.400000
0.400000
0.200000
Function Description

Complete the plusMinus function in the editor below.

plusMinus has the following parameter(s):

int arr[n]: an array of integers
Print
Print the ratios of positive, negative and zero values in the array. Each value should be printed on a separate line with  digits after the decimal. The function should not return a value.

Input Format

The first line contains an integer, , the size of the array.
The second line contains  space-separated integers that describe .
*/

package main

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	var positive, negative, zero, n float64
	n = float64(len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positive++
		}
		if arr[i] < 0 {
			negative++
		}
		if arr[i] == 0 {
			zero++
		}
	}

	fmt.Println(fmt.Sprintf("%.6f", float64(positive/n)))
	fmt.Println(fmt.Sprintf("%.6f", float64(negative/n)))
	fmt.Println(fmt.Sprintf("%.6f", float64(zero/n)))
}

func main() {
	var arr []int32 = []int32{1, 1, 0, -1, -1}
	plusMinus(arr)
}
