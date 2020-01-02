package main

import (
    "fmt"
)

func main() {
    // There are also the following variations of unsigned integers:
    // uint8 uint16 uint32 uint64 uintptr
    var unsignedIntegerDefault uint = 100
    fmt.Println("unsignedIntegerDefault = ", unsignedIntegerDefault)

    // Bytes are used quite frequently when dealing with strings in Go, and they are
    // provided as an alias for uint8 (0-255)
    var myByte byte = 65
    // Since 65 is 'A' in ASCII, we'll see it is easily represented as both a number and character
    fmt.Println("My byte as a number:", myByte, "My byte as a character:", string([]byte{myByte}))

    // Complex numbers have two parts -- the real and imaginary. In Go we represent them as a pair of either
    // `float32`s or `float64`s, depending on which complex datatype is used. See below:

    // complex64 allows `float32` real and imaginary parts
    var c64 complex64
    fmt.Println("c64 =", c64)

    // Also c64 by default
    c64Default := 21 + 42i
    fmt.Println("c64Default =", c64Default)

    // complex128 allows `float64` real and imaginary parts
    var c128 complex128
    fmt.Println("c128 =", c128)

    // Notice the syntax for the rune is a single-quoted character. This is specific to runes. You may not use ""
    var aRune rune = '🚀'
    // A rune is a single Unicode code-point and is actually just an alias for `int32`, so this will print a number
    fmt.Println("Rocket ship code-point =", aRune)
    // Since it's an integer, we can also just set it as such
    // 128640 == 🚀, whereas 11088 == ⭐
    aRune = 11088
    fmt.Println("Star code-point =", aRune)
}