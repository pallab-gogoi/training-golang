package main

import (
	"fmt"
	"os"
)

func main() {
	menue()
	var n int
	var code string
	var total []int
	no := 0
	priceCode := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	for true {
		fmt.Println("Plz place the order: ")
		fmt.Scan(&code)
		if code == "END" {
			var total_income int

			for i := 0; i < len(total); i++ {
				total_income = total[i] + total_income
			}
			fmt.Println("Total Income : ", total_income)
			os.Exit(0)
		}
		fmt.Scan(&n)
		bill := n * priceCode[code]
		fmt.Println("your bill is:", bill)

		total = append(total, bill)
		no = no + 1
	}
}
func menue() {
	fmt.Print("C: coffee, 40rs\n" +
		"D: dosa, 80 rs\n" +
		" T: tomato soup, 20rs\n" +
		"I : idli , 20rs\n" +
		"V : vada, 25rs\n" +
		"B: bhature &chhole, 30rs\n" +
		"P: paneer pakoda, 30rs\n" +
		"M: manchurian, 90rs\n" +
		"H: hakka noodle, 70rs\n" +
		"F: French fries, 30rs\n" +
		"J: jalebi, 30rs\n" +
		"L; lemonade, 15rs\n" +
		"S: spring roll, 20rs\n")
}
