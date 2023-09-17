package main

import (
	try "backend/try/list"
)

func main() {
	ll := new(try.LinkedList)

	ll.Insert(9)
	ll.Insert(2)
	ll.Insert(4)
	// ll.ReverseIterative()
	// ll.ReverseRecursive()
	ll.ReverseEven()
	ll.Display()

}

// package main

// import "fmt"

// type Employee struct {
// 	Number int
// 	Name   string
// }

// func main() {
// 	var N int
// 	fmt.Scan(&N)

// 	var employees []Employee
// 	for i := 0; i <= N-1; i++ {
// 		var Func string
// 		fmt.Scan(&Func)

// 		switch Func {
// 		case "make":
// 			var num int
// 			var name string
// 			fmt.Scan(&num)
// 			fmt.Scan(&name)
// 			employee := Employee{
// 				Number: num,
// 				Name:   name,
// 			}
// 			employees = append(employees, employee)
// 		case "getnum":
// 			var S int
// 			fmt.Scan(&S)
// 			value := getnum(S, employees)
// 			fmt.Println(value)
// 		case "getnamPrintlntry
// 			var S int
// 			fmt.Scan(&S)
// 			value := getname(S, employees)
// 			fmt.Println(value)
// 		}
// 	}
// }

// func getnum(num int, employees []Employee) int {
// 	return employees[num - 1].Number
// }

// func getname(name int, employees []Employee) string {
// 	return employees[name - 1].Name
// }
