package main

// Across these examples, we frequently use the built-in `len(something)` function,
// that gives us the number of elements in our array, map, or slice

func main() {
	// NB: Again, only look at the array-related stuff if you want to...
	// It's not really recommended or needed for beginning Go developers as it usually just confuses

	// The difference in the syntax compared to slices, is that array initializations/declarations will always have a
	// number (or ellipsis, if inferrred) in the brackets to specify length or to have the length inferred from the
	// list of values provided in the {} after the type

	// When declaring an array with this format, Go initializes each value to its type's default (0 in this case)
	var firstArray [10]int

	// When declaring an array literal, we can use the [...] syntax to have Go infer the length of the literal
	// For example, this array will have a length of 3
	anotherArray := [...]int{21, 42, -21}

	// NB: Since arrays are fixed length (the length is immutable), we cannot append to arrays,
	//     and you also cannot append an array to slice, so use arrays at your own risk... :)
}
