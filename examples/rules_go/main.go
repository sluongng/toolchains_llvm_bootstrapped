package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lgreet
// #include "greet.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Call the C greet function via cgo
	name := C.CString("Go Developer")
	defer C.free(unsafe.Pointer(name))
	
	greeting := C.greet(name)
	fmt.Println(C.GoString(greeting))
	
	// Call the C add_numbers function
	a := C.int(15)
	b := C.int(27)
	result := C.add_numbers(a, b)
	fmt.Printf("C says: %d + %d = %d\n", a, b, result)
	
	fmt.Println("\nSuccessfully called C functions from Go using cgo!")
	fmt.Println("This demonstrates rules_go working with the LLVM bootstrapped toolchain.")
}
