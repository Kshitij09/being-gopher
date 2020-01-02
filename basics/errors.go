package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// NB: Across these examples, we frequently use the built-in `len(something)` function,
//     that gives us the number of elements in our array, map, or slice

func main() {
	// Frequently, we'll be dealing directly with the errors like this
	err := doSomethingThatMightFail(43)
	if err != nil {
		fmt.Println("Oops! Something went wrong:", err)
	}
	fmt.Println("Wow everything worked!?")

	// Sometimes when delegating to a function wholly, we can let it handle the errors directly (this may often be common for top-level functions in your own code)
	printTheJSONIfPossible()
}

func doSomethingThatMightFail(someNum int) error {
	// Let's pretend we're actually doing something useful...
	if someNum != 42 {
		// This is how we create our own errors -- as you can see, the primary feature of an error object is the error message that we provide
		return errors.New("invalid number passed for checking")
	}
	return nil
}

func printTheJSONIfPossible() {
	// Don't worry if you haven't seen some of the syntax below before, it will be explained in upcoming sections
	jsonObj := map[string]string{}
	err := json.Unmarshal([]byte("This isn't valid JSON!!!"), &jsonObj)
	// Main point to note here is how a standard library function (json.Unmarshal) returns an error, we store it in a variable called, 'err'
	// and subsequently check if it's not equal to `nil`. If that's the case, then that means there was an error and we should do something (like notify the user)
	if err != nil {
		// We can also call `err.Error()` to get the error message as a string in our code -- or we can just let fmt print it out for us
		fmt.Println("There was an error parsing the JSON:", err)
		return
	}
	fmt.Println("Succesfully parsed the JSON:", jsonObj)
}
