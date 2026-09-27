//function signature - this is a decription of what a function does in terms of its input, output and its types

//a function that takes two strings and concats them together

package main

import "fmt"

func concat(a string , b string) string{
	return  a +b
}
//when two paramters are of the same type, stating the type of the last parameter is the same as stating the first

func concat2(a, b string) string{
	return  a +b
}

func main(){
	fmt.Println(concat("hello", "world"))
	fmt.Println(concat("hi", "people"))
}


//the defer keyword