package main

import (
	"fmt"
)

func main() {
	// We use the `var` keyword, followed by the variable name and type to declare a variable
	// NB: Go will also initialize the variable to its default value, which for an int is 0
	//     but for other types, that might be `nil`. So watch out! More on defaults in the "Datatypes Overview"
	var i int

	// We can assign a value to our variable
	i = 21

	// And we can re-assign as much as we want
	i = 42

	// NB: The type of a variable, once declared, cannot be changed
	//     e.g., this will cause an error `i = "Something that's not an int"`

	// NB: We must use a variable (not just declare and/or assign) it, otherwise Go will
	//     give us a compile-time error! In the case that you just want to let the compiler
	//     know that you're smarter than it, you can basically use your variable in a 'throw away'
	//     assignment to a special character in Go (the "blank identifier")

	// The `_` symbol in Go is not a variable and cannot be accessed, it's just a sentinel value that
	// we sometimes use to tell the compiler, "hey, I know what I'm doing" (it's a no-op,
	// like writing to /dev/null on a *nix system)
	_ = i

	// NB: the `_` can also be used to import a package for a side-effect (like having that package's `init`
	//     function executed). This is somewhat common, however, you shouldn't just use this import a ton of
	//     libs that you don't need as that will just slow down your compilation and increase the
	///    size of your executable...

	// The short-hand := operator (very commonly used) allows us to declare and assign the variable, while having Go
	// infer the type for us. This doesn't mean that it doesn't have a type!
	y := 4
	_ = y

	num, err := getSomeNumber()
	// `nil` is a keyword that's similar to `null` or `None` in other languages
	if err != nil {
		fmt.Println("Oops! We got an error:", err)
		return
	}
	fmt.Println("Successfully got our num:", num, "err (is <nil>):", err)

}

// Sample function to show how to assign multiple values from a function call
func getSomeNumber() (int, error) {
	return 42, nil
}
