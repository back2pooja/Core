package main

import (
	"fmt"
)

func main() {

	emp1 := Employee{Id: 001, name: "Pooja", Dept: "IT"}
	emp2 := Employee{Id: 002, name: "Priya", Dept: "Sales"}

	dept, err := emp1.GetDepartment()
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Dept", dept)
	}

	dept, err = emp2.GetDepartment()
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Dept", dept)
	}
}
