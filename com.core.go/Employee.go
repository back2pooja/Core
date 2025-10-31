package main

import "fmt"

type Employee struct {
	Id   int
	name string
	Dept string
}

func (e *Employee) GetDepartment() (string, error) {
	if e.Dept == "" {
		return "", fmt.Errorf("dept not found for the employee %s", e.name)
	}
	return e.Dept, nil
}
