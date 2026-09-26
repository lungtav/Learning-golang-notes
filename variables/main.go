//the basic variable types are

//string
//bool
//int - int8, int16, int32, int64
//uint - uint8, uint16, uint32, uint64
//float32, float64
//complex64, complex128
//byte - alias for uint8
//rune - alias for int32
//:= is used to declare variables without explicitly declaring the type

package main

import "fmt"

func main(){
	congrats:= "happy birthday"
	fmt.Println(congrats)


	//conditonals
	length:= 3
	if length < 1 {
		fmt.Println("Email is valid")
	}

	//the variable can also be initialized in the if /else block
	//the variable is then only accessible within the scope
	if length:= 3; length >= 3{
		fmt.Println("Email is valid")
	}
}