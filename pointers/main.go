package main

import "fmt"

// Pointer Receiver
type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	// 	A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored, not the actual data itself.

	// The * syntax defines a pointer:

	// var p *int

	myString := "hello"
	myStringPtr := &myString

	fmt.Println(myStringPtr)


// 	Dereference
// The * operator dereferences a pointer to get the original value.

// *myStringPtr = "world"                              // set myString through the pointer
// fmt.Printf("value of myString: %s\n", *myStringPtr) // read myString through the pointer
// // value of myString: world



	// pointer receiver
	// c := car{
	// 	color: "white",
	// }
	// c.setColor("blue")
	// fmt.Println(c.color)
	// // prints "blue"

	//values  to hwich pointer points is usually stored in the stack. it is stored in the heap only if the compiler cannot prove the variable i referencs before the function returns
	

}