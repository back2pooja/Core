package main

import (
	"fmt"

	"Core/com.core.golang"
)

func main() {

	s1 := com_core_golang.Student{RollNo: 01, Name: "Pooja", Marks: 80}
	s2 := com_core_golang.Student{RollNo: 21, Name: "Durga", Marks: 90}

	rest1, _ := s1.GetResults()
	rest2, _ := s2.GetResults()

	fmt.Println(s1.Name, "Result", rest1)
	fmt.Println(s2.Name, "Result", rest2)
}
