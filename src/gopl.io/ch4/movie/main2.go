package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	type Employee struct {
		Name   string
		Salary int
	}

	employees := []Employee{
		Employee{
			Name:   "Sarah",
			Salary: 30000,
		},
		Employee{
			Name:   "John",
			Salary: 25000,
		},
	}

	data, err := json.MarshalIndent(employees, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	err = ioutil.WriteFile("employees.json", data, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed writing file: %s", err)
	}
}
