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

func Example_rangePointers() {
	ints := []int{1, 2, 3}
	ptrs := []*int{}
	// golang uses a single pointer when computing a range,
	// so this produces a slice of the same length as the
	// original slice, but filled with the same pointer
	for _, i := range ints {
		ptrs = append(ptrs, &i)
	}

	// since the range pointer is now pointing to the last
	// value in the slice, we print the same value three times
	for _, ptr := range ptrs {
		fmt.Println(*ptr)
	}

	// Output:
	// 3
	// 3
	// 3
}
