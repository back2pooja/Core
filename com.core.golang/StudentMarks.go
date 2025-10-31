package main

import (
	"fmt"
)

type Student struct {
	RollNo int
	Name   string
	Marks  int
}

func (s *Student) GetResults() (string, error) {
	if s.Marks < 0 {
		return "", fmt.Errorf("invalid Marks for student %s", s.Name)
	}
	if s.Marks >= 40 {
		return "Pass", nil
	}
	return "Fail", nil
}

func main() {

	s1 := Student{RollNo: 01, Name: "Pooja", Marks: 80}
	s2 := Student{RollNo: 21, Name: "Durga", Marks: 90}

	rest1, _ := s1.GetResults()
	rest2, _ := s2.GetResults()

	fmt.Println(s1.Name, "Result", rest1)
	fmt.Println(s2.Name, "Result", rest2)
}
