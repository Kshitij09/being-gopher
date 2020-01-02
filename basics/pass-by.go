package main

import (
	"fmt"
)

type myFoo struct {
	amount int
}

func main() {
	// Setup our demo variables
	mySliceOfNums := []int{0, 42, 21, -1}
	myStruct := MyFoo{amount: 42}

	fmt.Println("mySliceOfNums before call to changeMySlice:", mySliceOfNums)
	// Let the slice be changed!
	changeMySlice(mySliceOfNums)
	fmt.Println("mySliceOfNums after call to changeMySlice:", mySliceOfNums)
	// Can't change the reference to the slice I pass...
	youCantTouchThisSlice(mySliceOfNums)
	fmt.Println("mySliceOfNums after call to youCantTouchThisSlice:", mySliceOfNums)

	fmt.Println("myStruct before call to youCantChangeThisFoo:", myStruct)
	youCantChangeThisFoo(myStruct)
	fmt.Println("myStruct after call to youCantChangeThisFoo:", myStruct)
	// The & is used to pass a pointer to myStruct
	changeFooAmount(&myStruct)
	fmt.Println("myStruct after call to changeFooAmount:", myStruct)
}

func changeMySlice(numbers []int) {
	if len(numbers) > 0 {
		numbers[0] = 42
	}
}

func youCantTouchThisSlice(numbers []int) {
	// I'm going to try to change the slice itself, but I can't because I'm actually just changing
	// a reference to the slice -- as opposed to the slice itself -- which is passed by value
	numbers = []int{}
}

func youCantChangeThisFoo(foo MyFoo) {
	// So this will only change my copy of `foo` -- and the caller will not see any changes
	foo.amount = -1
}

func changeFooAmount(refToFoo *MyFoo) {
	// So, I can change a field in the refToFoo, because I'm modifying the actual value that my pointer points to
	// and the caller will subsequently see their `amount` as 21
	refToFoo.amount = 21
	// However, if I just change the pointer, nothing will happen and the caller will still see `amount` as 21...
	refToFoo = &MyFoo{amount: 0}
}
