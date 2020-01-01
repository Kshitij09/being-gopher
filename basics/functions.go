package main

import (
	"fmt"
	"strings"
)

// We use this syntax to define a function in Go
// The function `main()` is reserved for the entry-point of our program
func main() {
	// Invoke the sayHelloAndReturnNothingness function without any parameters
	sayHelloAndReturnNothingness()

	// Call getMyNumber() and print out the response using fmt
	fmt.Printf("getMyNumber() -> %d\n", getMyNumber())

	// Use the combined type inferrence, declaration, and assignment operators to get the results of the getMyFavoriteNumber func
	myFaveNum, reason := getMyFavoriteNumber()

	// NB: Remember that fmt.Println will put spaces between arguments for us!
	fmt.Println("My Favorite Number:", myFaveNum, "For The Reason:", reason)

	// Let's pass a function to another function as an argument
	fmt.Println(giveMeSomeFunction(getMyNumber))

	// Show me a use case for passing around functions!
	// NB: the `:=` operator is a short-hand way of declaring and defining variables, more on this in the Variables section!
	returnedFunc := funcFactory("Show", "how", "much", "cool.")
	fmt.Println(returnedFunc("Wow", "such", "factory."))
	fmt.Println(returnedFunc("See!", "There", "is", "a", "use", "case!"))

	// Join and append
	arr := []string{"a", "b", "c", "d"}
	fmt.Println(strings.Join(arr, ","))
	fmt.Println(append(arr, "x", "y", "z"))
}

// This straightforward syntax defines a *package-wide* function that doesn't return anything
func sayHelloAndReturnNothingness() {
	fmt.Println("I'm a F U N C T I O N in Go! Cool.")
}

// Let's return an `int` from this function
func getMyNumber() int {
	return 42
}

// Let's return more than one value from a function (OMG, we can really do that? You bet!)
// NB: whenever using 'named' return values, they are (1) automatically defined with their default values and;
//     (2) we must use parenthesis around them to let Go know what we're up to!
func getMyFavoriteNumber() (num int, reason string) {
	// Since Go defined `num` and `reason` from the function signature above, we can use the assignment operator `=`
	num = 42
	reason = "#TheAnswerToLife"
	// Below, Go implicitly returns `num` and `reason`, however, the following would also be equivalent:
	// `return num, reason` which in this case would also be equivalent to:
	// `return 42, "#TheAnswerToLife"`
	return
}

func giveMeSomeFunction(f func() int) int {
	return f()
}

// By the way, this is also a variadic function, you can check that out in the extras in a section called "Variadic Functions"
func funcFactory(prefix ...string) func(suffix ...string) string {
	return func(suffix ...string) string {
		return strings.Join(append(prefix, suffix...), " ")
	}
}
