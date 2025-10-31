package com_core_golang

import "fmt"

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
