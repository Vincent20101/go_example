package main

import (
	"fmt"
	"reflect"
)

type MeasurementMethod struct {
	Event bool `json:"event,omitempty"`
	Volum bool `json:"volum,omitempty"`
	Durat bool `json:"durat,omitempty"`
}

func main() {
	var measurementMethod MeasurementMethod
	measurementMethod.Volum = true
	if measurementMethod == (MeasurementMethod{}) {
		fmt.Println("Measurement", measurementMethod)
	}

	fmt.Println(IsZeroValue(measurementMethod))
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

func IsZeroValue(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}