// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 1

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)

	// print the value of the package-level variable "val"
	printPackageVar()
	// update the package-level variable "val" and print it again
	updatePackageVar()
	printPackageVar()
	// print the pointer address of the local variable "val"
	printPackageVarAddr()
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 100
	fmt.Printf("local val = %v\n", val)
}

func printPackageVar() {
	fmt.Printf("%T, golbal val = %v\n", val, val)
}

func printPackageVarAddr() {
	fmt.Printf("%T, local &val = %v\n", &val, &val)
}

func updatePackageVar() {
	val = "updated global"
}
