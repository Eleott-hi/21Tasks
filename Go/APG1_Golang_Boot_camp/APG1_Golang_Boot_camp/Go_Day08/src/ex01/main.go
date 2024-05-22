package ex01

import (
	"fmt"
	"reflect"
	"strings"
)

func DescribePlant(plant interface{}) {
	// Get the type of the plant
	t := reflect.TypeOf(plant)
	if t.Kind() != reflect.Struct {
		fmt.Println("Expected struct")
		return
	}

	// Get the value of the plant
	v := reflect.ValueOf(plant)
	if t.NumField() == 0 {
		fmt.Println("Empty struct")
		return
	}

	// Iterate through the fields of the plant
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if field.Tag == "" {
			fmt.Printf("%s:%v\n", field.Name, value)
		} else {
			tag := strings.Fields(string(field.Tag))[0]
			pairs := strings.Split(tag, ":")
			key := pairs[0]
			val := field.Tag.Get(key)
			fmt.Printf("%s(%s=%s):%v\n", field.Name, key, val, value)
		}
	}
}
