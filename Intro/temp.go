package main

import "fmt"

func main() {
	//var x float64
	var arr []int32 = []int32{4, 4, 1, 3, 99, 8, 7, 6, 1, 2, 4, 32, 3, 52, 3, 4, 2, 4, 6, 2, 43, 3, 2, 4, 23, 5, 2, 23, 5, 3, 2, 4, 5, 2, 3, 5}
	fmt.Println(birthdayCakeCandles(arr))
}
func largest(arr []int32) int32 {
	max := arr[0]

	for i := 1; i < len(arr); i++ {

		if max < arr[i] {

			max = arr[i]
		}
	}
	return max

}
func birthdayCakeCandles(candles []int32) int32 {
	large := largest(candles)
	fmt.Println(large)
	var count int32
	for i := 0; i < len(candles); i++ {
		if large == candles[i] {
			count++
		}
	}
	return count
}
