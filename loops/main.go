package main

import "fmt"

func main() {

	for i := 0; i < 30; i++ {
		fmt.Println(i)
	}

	//it can also be written as
	for i:= range 30{
		fmt.Println(i)
	}

	//conditions can be removed which makes it run forever
	// for i:= 0; ;i++{
	// 	fmt.Println(i)
	// }

	//since there no while loop in go
	//a for can be used

	for 1 < 3{
		fmt.Println("hello")
	}
}
