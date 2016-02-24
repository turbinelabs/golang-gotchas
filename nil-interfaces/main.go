package main

import (
	"fmt"
)

type foo struct {
	Field1 *int    `json:"field1"`
	Field2 *string `json:"field2"`
}

func isInterfaceNil(i interface{}) {
	fmt.Printf("interface %s == nil? %v\n", i, i == nil)
}

func main() {
	var i interface{}

	var f *foo

	fmt.Println(i)        // i is <nil, nil>; it has no type and no value
	fmt.Println(f)        // f is a staticly typed *foo with value nil; it can be thought of as <*foo, nil>
	fmt.Println(i == nil) // i == nil because "interface{} nil" is <nil, nil> which is equal to the zero value of i
	fmt.Println(f == nil) // f == nil because "*foo typed nil" is <*foo, nil> which matches the zero value of f

	// It gets a bit weird here though
	fmt.Println(i == f) // i (<nil, nil>) != f (<*foo, nil>) even though i and f both "== nil"...
	// This is because the zero value that nil is expanded to for the comparison
	// to i is not the same zero value used when checking f. This means that if
	// we compare i (<nil, nil) to f (<*foo, nil>) they are not equal.

	fmt.Println()
	fmt.Println()
	// in practice you can think of all types as having a dynamic type equal to their
	// actual type and an interface as having a mutable dynamic type. But if the mutable
	// dynamic type is set then the zero-value (nil) of an interface{} is no longer
	// equal because nil implies an unset dynamic type. in the non-interface case
	// the zero-value (nil) implies the appropriately fixed "dynamic" type.

	// so if we assign the zero-valued *foo to i you mutate the dynamic type value of
	// i so that i == f is now true but i == nil is not
	i = f
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(i == f)
	fmt.Println(f == nil)
	fmt.Println(i == nil)

	fmt.Println()
	fmt.Println()

	// all together: if you pass a type into an interface{} fuck everything:
	var i2 interface{}
	var f2 *foo

	isInterfaceNil(nil)
	isInterfaceNil(i2)
	isInterfaceNil(f2)

	// and this is why using typed errors in go is a bad smell (except in places where
	// you have a very good understanding of the lifecycle). Which is, to draw inspiration
	// from Dana, a garbage fire.
}
