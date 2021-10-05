package main

import "fmt"

func main() {
	var a uint8 = 5 				// Unsigned 8-bit integer, 0 to 255
	var b uint16 = 60000 			// Unsigned 16-bit integer, 0 to 65535
	var c uint32 = 2000000 			// Unsigned 32-bit integer, 0 to 4294967295
	var d uint32 = 40889999			// Unsigned 64-bit integer, 0 to 18446744073709551615
	fmt.Println(a, b, c, d)

	var e int8 = -5					// Singed 8-bit integer, -128 to 127
	var f int16 = 1257				// Signed 16-bit integer, -32768 to 32767
	var g int32 = 6545485			// Signed 32-bit integer, -2147483648 to 2147483647
	var h int64 = 40889999			// Signed 64-bit integer, -9223372036854775808 to 9223372036854775807
	fmt.Println(e, f, g, h)

	var i float32 = 32.52			// 32-bit floating number
	var j float64 = .587878			// 64-bit floating number
	fmt.Println(i, j)

	var k complex64 = 2i + 5 		// 64-bit complex number
	var l complex128 = 2123i + .5 	// 128-bit complex number
	fmt.Println(k, l)

	var m byte = 5					// Same as uint8
	var n rune = 58848				// Same as int32
	var o uint = 2545				// 32 or 64 bit
	var p int = 5878				// 32 or 64 bit
	var q uintptr = 12312			// Uint for pointers
	fmt.Println(m, n, o, p ,q)

	var r bool = true				// Boolean type, true or false
	var s string = "hey"			// String type
	fmt.Println(r, s)
}
