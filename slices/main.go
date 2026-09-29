package main

import "fmt"

func main() {

	//arrays in go
	//to declare an  array of 10 integers
	var myInts [10]int
	fmt.Println(myInts)

	primes := []int{2, 3, 5, 7, 11, 13}
	newPrimes := append(primes, 1)
	fmt.Println(newPrimes)

	mySlice := primes[1:4]
	fmt.Println(mySlice)
	mySlice = append(mySlice, 1)
	fmt.Println(mySlice)
	// the zero value of slice is nil

	// func make([]T, len, cap) []T
	// mySlice := make([]int, 5, 10)

	// the capacity argument is usually omitted and defaults to the length
	// mySlice := make([]int, 5)

	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}
