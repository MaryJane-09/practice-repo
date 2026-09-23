package main

import "fmt"

func main() {
	// Primitive & overflow
	fmt.Println("Primitive & overflow")
	var num int8 = 127

	fmt.Println(num+1) //output = -128
	// this happens as a result of overflowing there is nothing like 128 in int8 so 127(the last number in int8) +1 overflows and automatically counts back
	fmt.Println()

	// Mutability
	fmt.Println("Mutability")
	var slice = []int{4, 3, 7}
	updated := append(slice, 10)
	fmt.Println(updated) // output = [4 3 7 10] it changed
	fmt.Println(slice) //output = [4 3 7]
	fmt.Println()

	//Value vs reference — struct/array
	fmt.Println("Value vs reference — struct/array")
	type food struct{rice, beans, yam string}
	newfood := food{"jellof", "porriage", "fried"}
	new := newfood
	new.yam = "boiled"
	fmt.Println(newfood.yam) //output = fried
	fmt.Println()

	// pointer
	fmt.Println("Pointers")
	age := 17 // variable with a value

	fmt.Println(age) // printing the value

	fmt.Println(&age) // printing address of the value

	ptr := &age       // when you assign the address of a value to another variable
	fmt.Println(*ptr) // and print it using '*' it prints pack the original value
	fmt.Println(ptr)  // else it still prints the address of the value
	*ptr = 25        // use a pointer operator to change the value of the variable and since 'ptr' already points to 'age' it the new variable will be printed
	fmt.Println(age) //Changing what the pointer points to changes the original variable.

	
	
}
