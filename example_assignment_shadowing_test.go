/*
Copyright 2017 Turbine Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package golanggotchas_test

import (
	"fmt"
)

func getValue(r int) (int, error) {
	return r, nil
}

func display(i int, e error) {
	fmt.Printf("(%v, %v)\n", i, e)
}

func Example_assignmentShadowing() {
	// Start out with a nice assignment; since we want combined declaration &
	// assignment use :=
	foo := 1000
	display(foo, nil) // print out the value; 1000 of course

	if true {
		// call a function that returns (int, error); since err isn't declared yet
		// we use := for mixed declared and undeclared variables
		foo, err := getValue(2000)
		display(foo, err) // print out value; 2000 as expected
	}

	// lets check back in on foo for giggles
	display(foo, nil) // wait... it is now 1000 again.

	fmt.Println("----")

	// Okay. What's happening here is that := prefers to declare over assign
	// variables. That means if you try to use it in a block it's going to shadow
	// all your external variables.

	// In order to maintain your previously declared variables you need to
	// explicitly declare the block-scoped variables:

	display(foo, nil) // still starting at 1000
	if true {
		var err error             // declare within the block
		foo, err = getValue(3000) // now this is only assignment
		display(foo, err)         // displays the expected 3000
	}

	// And now, because we didn't use := above, foo retains 3000
	display(foo, nil)

	fmt.Println("----")

	// This is surprising as typical, same-scope, use of := for declaration &
	// assignment will actually have the expected behavior with the existing
	// variable assigned and the new variable declared. This behavior is shown
	// below.

	baz := 1000 // declare a new variable
	display(baz, nil)

	baz, err := getValue(4000) // use := to for mixed declaration & assignment
	display(baz, err)          // and baz is updated as expected

	// Output:
	// (1000, <nil>)
	// (2000, <nil>)
	// (1000, <nil>)
	// ----
	// (1000, <nil>)
	// (3000, <nil>)
	// (3000, <nil>)
	// ----
	// (1000, <nil>)
	// (4000, <nil>)
}
