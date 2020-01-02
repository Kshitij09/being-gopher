package main

import (
	"fmt"
)

func main() {
	showIfStatements()
	showSwitchStatements()
	showForStatements()
	fmt.Println("Got", showGotoStatements(), "from showGotoStatements()")
}

func showIfStatements() {
	t, f := true, false
	// Note that `if` statements don't have parentheses, only braces around the code block
	if true {
		fmt.Println("True is definitely true!")
	}

	if t == f {
		fmt.Println("Nope!")
	} else if t != f {
		fmt.Println("Makes sense")
	} else if t == !f {
		fmt.Println("Also true, but we didn't get this far")
	} else {
		fmt.Println("Nope ...")
	}
}

func showSwitchStatements() {
	// Sometimes it's best to use switch in order to avoid long chains of if/else if/else's
	myNum := 21
	switch myNum {
	case 1, 2, 42:
		// This would get called if myNum == 1 OR 2 OR 42
		fmt.Println("Cool. myNum is either 1, 2, or 42")
		// ⚠️ Remember that in Go, cases DO NOT fall through by default!
		break // So this is actually redundant here, but just to show that we _can_ manually break as well
	case 3:
		// This is a no-op case, i.e., if we _want_ to do nothing when myNum == 3
	case 5:
		// ⭐ If you want to fall through, you can (the fallthrough must be the last line of that case)
		fallthrough
	default:
		// Optionally, we can have a default case that gets applied either via an explicit fallthrough or when no case matches
	}
}

func showForStatements() {
	// Setup some dummy data (these will be explained in "Core Datatypes" so don't worry if they are unfamiliar right now)
	colorsMap := map[string]string{
		"AliceBlue":    "#F0F8FF",
		"AntiqueWhite": "#FAEBD7",
		"Aqua":         "#00FFFF",
		"Aquamarine":   "#7FFFD4",
		// NB: Go requires this trailing comma on the last element when initializing values over multiple lines
		"Azure": "#F0FFFF",
	}

	colorsSlice := make([]string, 0, len(colorsMap))

	// Here, we use the `range` keyword to have Go iterate over the `colorsMap` and set `name` to the "key" in
	// the map, while setting `hexStrValue` to the value for that particular key in the map.
	for name, hexStrValue := range colorsMap {
		colorsSlice = append(colorsSlice, fmt.Sprintf("%s:%s", name, hexStrValue))
	}

	// You can also use range on slices (as below) -- and even `chan`s (explained further in concurrency)
	// If you don't need the key or index, you can use the blank `_` to ignore it
	for _, colorPairValue := range colorsSlice {
		fmt.Println(colorPairValue)
	}

	// NB: Like `if` statements, `for` loops do not use parenthesis -- just braces around the inner block

	// Iteration in Go is essentially the same as in Java/C/C++
	for x := 0; x < 3; x++ {
		// Remember that variables declared in `for` and `if` statements or blocks are scoped to that block
		fmt.Println("iteration", x)
	}

	// Since we don't have while loops, we can also use this form of `for` to create infinite loops
	for {
		// `break` allows us to leave the for loop
		break
		// `continue` allows us to move on to the next iteration in the loop (leaving the current iteration)
		continue // This wouldn't be reached in this case, as we used `break` to leave above
	}
}

func showGotoStatements() int {
	// This example isn't particularly useful (for clarify and brevity)
	fmt.Println("Let's see how goto works...")
	i := 0
	if i == 0 {
		goto namedBlock
	} else {
		goto otherBlock
	}
	// That's it, the goto will make one of these blocks finsh and return from the function

passBack:
	return i

namedBlock:

	fmt.Println("Hey, I was invoked via a goto statement!")
	i++
	goto passBack

otherBlock:

	fmt.Println("Hey, I'm another block. Did I get invoked??")
	i--
	goto passBack
}
