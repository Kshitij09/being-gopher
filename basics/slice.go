package main

import (
	"fmt"
)

// NB: Across these examples, we frequently use the built-in `len(something)` function,
//     that gives us the number of elements in our array, map, or slice

func main() {
	// We use `make` to create an empty slice of our desired type, length, and, optionally, capacity
	// NB: Below, 0 is the length of our slice, and 10 is the capacity. This means that Go
	//     will allocate memory for 10 elements and `append()` will be a cheap operation until we exceed
	//     that capacity (where the Go runtime would automagically expand the slice for us -- expensive though!)
	// NB: `make([]string, 10)` would create a slice with 10 elements (each of which would be initialized
	//     to their default value based on their type) and the capacity would be equal to the length.
	demoSlice := make([]string, 0, 10)

	// Slice will be [] and len will be 0
	fmt.Println("Empty slice:", demoSlice, "with len:", len(demoSlice))

	// We can also create slice literals, like so
	demoSlice = make([]string, 10)

	// Slice with defaulted elements
	fmt.Println("Slice filled with defaults:", demoSlice, "with len:", len(demoSlice))

	// Set values for specific indices
	demoSlice[0] = "21"
	demoSlice[9] = "42"

	// If you use an index, like 10 in this case, that is out of bounds, you will get a panic (similar
	// to an exception in some other languages)
	// e.g., `demoSlice[10]` would cause a panic, and in this case cause the program to exit with a failure status code

	fmt.Println("Get a slice of a slice:", demoSlice[:3])
	fmt.Println("Another slice of a slice:", demoSlice[3:])
	fmt.Println("Yet another slice of a slice:", demoSlice[2:6])
	fmt.Println("Valid, yet useless, slice of a slice:", demoSlice[:])

	letsAppendSlice := []string{"foo", "appendable"}
	// The append function will not modify the slices, so if you want to update the slice with the appended version, then you must also assign the result
	// NB: the letsAppendSlice... notation just expands that slice to send to the variadic append function
	demoSlice = append(demoSlice, letsAppendSlice...)
	// Would've also been equivalent to: `demoSlice = append(demoSlice, "foo", "appendable")`
	// demoSlice is now larger, letsAppendSlice is unchanged
	fmt.Println("Updated slice:", demoSlice)
}
