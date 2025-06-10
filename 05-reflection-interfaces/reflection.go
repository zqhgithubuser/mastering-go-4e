package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{
		"String value",
		-12.123,
		Secret{"Mihalis", "Tsoukalos"},
	}

	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String()) // <main.Record Value>

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)                            // main.Record
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType) // The 3 fields of main.Record are
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type()) // main.Secret
		}
		// 同上
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type()) // main.Secret
		}
	}
}
