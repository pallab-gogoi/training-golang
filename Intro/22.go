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
