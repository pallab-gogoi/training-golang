// interface: a data type, abstract data type
// interface have only method singnature, we can have empty interface also
// interface can hold variable data types
/*
type interface_name interface{
	method-1
	method-2
}
*/
package main

import (
	"fmt"
)

type Integer int
type number int

func (i Integer) print() {
	fmt.Println(i)
}
func (i Integer) CalculateSalary() int {
	return int(i)
}

func main() {

	j := Integer(1009)
	j.print()
	// var a int
	// a = 10
	// a = "10"

	// fmt.Print(a)
	// pj := PermanentJob{
	// 	basicpay: 10,
	// }
	//pj.CalculateSalary()//10
	// var sc SalaryCalculator
	// sc = pj
	// total := sc.CalculateSalary()
	// fmt.Println(total)

	pj := PermanentJob{
		basicpay: 10,
	}
	cj := ContractJob{
		basicpay: 20,
	}
	fj := FreelanceJob{
		basicpay: 1000,
	}
	i := Integer(222)

	sc := []SalaryCalculator{pj, cj, fj, i}
	// // total := sc.CalculateSalary()
	// // fmt.Println(total)
	// // sc.GetJobName()
	totalIncome(sc)

}
func totalIncome(sc []SalaryCalculator) {
	total := 0
	for _, v := range sc {
		total = total + v.CalculateSalary()
	}
	fmt.Println(total)
}

//interface - data type - abstract data type
//interface have only methods signature, well we can have empty interface also
/*
type interface_name interface{
	method-1
	method-2
}
*/
//
type SalaryCalculator interface {
	CalculateSalary() int
}
type PermanentJob struct {
	basicpay int
}

func (p PermanentJob) CalculateSalary() int {
	return p.basicpay
}

func (p PermanentJob) GetJobName() string {
	return "Permanet Job"
}

type ContractJob struct {
	basicpay int
}

func (c ContractJob) CalculateSalary() int {
	return c.basicpay
}

type FreelanceJob struct {
	basicpay int
}

func (f FreelanceJob) CalculateSalary() int {
	return f.basicpay
}

// var sc  SalaryCalculator
// this sc can hold the data of multiple data types
// which comes with a certain condition -
// sc should hold the variable of type PermanentJob and ContractJob
// the condition is that types(i.e PermanentJob and ContractJob ) must implement all the methods of interface
// interface will group the different data types into a single data type
