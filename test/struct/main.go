package main

import "fmt"

type MeasurementMethod struct {
	Event bool `json:"event,omitempty"`
	Volum bool `json:"volum,omitempty"`
	Durat bool `json:"durat,omitempty"`
}

func main() {
	var measurementMethod MeasurementMethod
	if measurementMethod == (MeasurementMethod{}) {
		fmt.Println("Measurement", measurementMethod)
	}

	type stu struct {
		age  int
		name string
	}

	var s stu
	s1 := stu{
		age:  10,
		name: "",
	}
	if s == s1 {
		fmt.Println("equal")
	}

}
